# Simple backlight adjust and low power mode

Designed and tested for raspberry pi with official 7 inch touchscreen. 

<br>

Other scripts that do the same thing:
 - C: https://github.com/WestleyK/rpi-brightness
 - Perl: https://github.com/WestleyK/pi-backlight
 - Bash/Shell: https://github.com/WestleyK/backlight

<br>

## Install, Update, and Uninstall:

<br>

Install via `go get`:

```
go get -u github.com/WestleyK/rpi-backlight
```

Or, clone the source code:

```
cd /tmp
git clone https://github.com/WestleyK/rpi-backlight
cd rpi-backlight/
make
sudo make install
```
<br>

### Should you install the pre-compiled code?

**Why install the pre-compiled:**
 - No dependency needed to install, (not that there is any).
 - You don't need golang to compile the source code.

**Why not to install the pre-compiled code:**
 - May not be fully compatible, (could only be partially compatible)
<br>

### Install with one line:

Copy-paste this in your terminal:
```
wget https://raw.githubusercontent.com/WestleyK/rpi-backlight/master/pre-compiled/armv7l/rpi-backlight
```
This will install `rpi-backlight` to the current directory.

To finish the install, do:
```
chmod +x rpi-backlight
sudo mv rpi-backlight /usr/local/bin/
```
<br>

### Uninstall:

```
cd ~/raspberrypi-backlight/rpi-backlight/
sudo make uninstall
```

Or uninstall manualy:

```
sudo rm /usr/local/bin/rpi-backlight
```

## Usage:

```
pi@raspberrypi:~ $ rpi-backlight --help
Usage: rpi-backlight [OPTION]
      --help         : print help menu.
      [15-255]       : adjust from: 15 to: 255.
      -s, --sleep    : enter sleep mode, press <ENTER> to exit this mode.
      -u, --up       : adjust brightness up by: 25/255.
      -d, --down     : adjust brightness down by: 25/255.
      -c, --current  : print current brightness.
      -n, --on       : turn backlight on to: 180.
      --version      : print script version.

Copyright (c) 2018-2019 WestleyK, All rights reserved.
This software is licensed under a Clear BSD License.
Source code: https://github.com/WestleyK/rpi-backlight
pi@raspberrypi:~ $ 
```

<br>
