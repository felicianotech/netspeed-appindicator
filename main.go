package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/gopherlibs/appindicator/appindicator"
	"github.com/gotk3/gotk3/gtk"
	"github.com/showwin/speedtest-go/speedtest"
)

func getSpeed() *speedtest.Server {

	user, _ := speedtest.FetchUserInfo()
	serverList, _ := speedtest.FetchServerList(user)
	targets, _ := serverList.FindServer([]int{})

	targets[0].PingTest()
	targets[0].DownloadTest(false)
	targets[0].UploadTest(false)

	return targets[0]
}

func main() {

	gtk.Init(nil)

	menu, err := gtk.MenuNew()
	if err != nil {
		log.Fatal(err)
	}

	indicator := appindicator.NewWithPath("indicator-netspeed", "netspeed-icon", appindicator.CategoryApplicationStatus, "/home/felicianotech/Repos/felicianotech/netspeed-appindicator")
	indicator.SetStatus(appindicator.StatusActive)
	indicator.SetMenu(menu)
	indicator.SetLabel(" n/a MB/s", "")

	itemStatus, err := gtk.MenuItemNewWithLabel("Speed: unknown")
	if err != nil {
		log.Fatal(err)
	}

	sep1, err := gtk.SeparatorMenuItemNew()
	if err != nil {
		log.Fatal(err)
	}

	itemWebsite, err := gtk.MenuItemNewWithLabel("Open Speedtest.net")
	if err != nil {
		log.Fatal(err)
	}
	_ = itemWebsite.Connect("activate", func() {
		exec.Command("xdg-open", "https://speedtest.net").Start()
	})
	if err != nil {
		log.Fatal(err)
	}

	sep2, err := gtk.SeparatorMenuItemNew()
	if err != nil {
		log.Fatal(err)
	}

	itemExit, err := gtk.MenuItemNewWithLabel("Quit")
	if err != nil {
		log.Fatal(err)
	}
	_ = itemExit.Connect("activate", func() {
		gtk.MainQuit()
	})
	if err != nil {
		log.Fatal(err)
	}

	menu.Add(itemStatus)
	menu.Add(sep1)
	menu.Add(itemWebsite)
	menu.Add(sep2)
	menu.Add(itemExit)
	menu.ShowAll()

	go func() {

		for {

			s := getSpeed()

			label := fmt.Sprintf("↓%6.1fMbps ↑%6.1fMbps\nlatency: %s", s.DLSpeed, s.ULSpeed, s.Latency.Round(time.Millisecond))
			mainLabel := fmt.Sprintf(" ↓%6.2fMbps", s.DLSpeed)

			indicator.SetLabel(mainLabel, "")
			itemStatus.SetLabel(label)

			<-time.After(time.Minute * 15)
		}
	}()

	gtk.Main()
}
