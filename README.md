# Simple backlight adjust and low power mode

<br>
<br>



## Install, Update and Uninstall:

<br>

### Install:

```
cd ~/
mkdir raspberrypi-backlight
cd raspberrypi-backlight/
git clone https://github.com/WestleyK/rpi-backlight
cd rpi-backlight/
make
sudo make install
```
<br>

### Installing pre-compiled code:

Why install pre-compiled:
 - No dependency needed to install, (not that there is any).
 - You don't need golang to compile the source code.

Why not to install pre-compiled code:
 - May not be fully compatible, (could only be partially compatible)
 - And many more reasons!
<br>

```
cd ~/
mkdir raspberrypi-backlight
cd raspberrypi-backlight/
git clone https://github.com/WestleyK/rpi-backlight
cd rpi-backlight/pre-compiled/
sudo ./make.sh install
```

### Install with one line:

Paste this in the terminal:
```
curl https://raw.githubusercontent.com/WestleyK/rpi-backlight/master/one-line-install.sh | sh
```
To finish the install, do:
```
sudo mv rpi-backlight /usr/local/bin/
```

**NOTE:** this will install the pre-compiled source code.

<br>
<br>

### Update:

```
cd ~/raspberrypi-backlight/rpi-backlight/
make update
sudo make install
```

<br>

### Uninstall:

```
cd ~/raspberrypi-backlight/rpi-backlight/
make uninstall
```
Or uninstall manualy:
```
sudo rm /usr/local/bin/rpi-backlight
```

<br>
<br>

## Usage:

```
pi@raspberrypi:~ $ rpi-backlight -help
Usage: rpi-backlight [OPTION]
      -help | --help (print help menu)
      -s | -sleep (enter sleep mode, press <ENTER> to exit this mode)
      -u | -up (adjust brightness up by: 25/255)
      -d | -down (adjust brightness down by: 25/255)
      -c | -current (print current brightness)
      -n | -on (turn backlight on to: 180
      -i | -info (print info)
      -version | --version (print version)
Source code: https://github.com/WestleyK/rpi-backlight
pi@raspberrypi:~ $ 
```


### End README


<br>
<br>



