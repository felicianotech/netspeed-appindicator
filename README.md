# NetSpeed AppIndicator

[![Build Status](https://circleci.com/gh/felicianotech/netspeed-appindicator.svg?style=shield)](https://app.circleci.com/github/felicianotech/netspeed-appindicator) [![Software License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/felicianotech/netspeed-appindicator/master/LICENSE)

NetSpeed AppIndicator is an Ubuntu AppIndicator for easy to access and frequent SpeedTest.net tests.


## Installing

### Debian Package (.deb) Instructions

Download the `.deb` file to the desired system.

Via your browser you can download it from the [GitHub Releases page][gh-releases].
You can then double-click the file to install.

Via terminal, you can do the following:

```bash
wget https://github.com/felicianotech/netspeed-appindicator/releases/download/v0.1.0/netspeed-appindicator-v0.1.0-amd64.deb
sudo dpkg -i netspeed-appindicator-v0.1.0-amd64.deb
```

`0.1.0` and `amd64` may need to be replaced with your desired version and CPU architecture respectively.


## Starting

You can start the NetSpeed AppIndicator by running the following:

```bash
systemctl --user start netspeed-appindicator
```

and keep it running between reboots:

```bash
systemctl --user enable netspeed-appindicator
```


## Updates

For now, you'll have to manually update by downloading new .debs and installing.
Watch this repo's releases to stay informed or just check back frequently.

## License

This repository is licensed under the MIT license.
The license can be found [here](./LICENSE).



[gh-releases]: https://github.com/felicianotech/netspeed-appindicator/releases
