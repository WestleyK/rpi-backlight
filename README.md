# Simple backlight adjust and low power mode

Designed and tested for raspberry pi with official 7 inch touchscreen. 

<br>

Other script that do the same thing:
 - c: https://github.com/WestleyK/rpi-brightness
 - perl: https://github.com/WestleyK/pi-backlight
 - bash/shell: https://github.com/WestleyK/backlight

<br>

<br>


## Install, Update and Uninstall:

<br>

### Install the source code:

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

**Why install pre-compiled:**
 - No dependency needed to install, (not that there is any).
 - You don't need golang to compile the source code.

**Why not to install pre-compiled code:**
 - May not be fully compatible, (could only be partially compatible)
<br>

### Install with one line:

Paste this in your terminal:
```
curl https://raw.githubusercontent.com/WestleyK/rpi-backlight/master/easy-install.sh | sh
```
This will install `rpi-backlight` to the current directory.

To finish the install, do:
```
sudo mv rpi-backlight /usr/local/bin/
```
<br>

#### Install by cloning repo:

```
cd ~/
mkdir raspberrypi-backlight
cd raspberrypi-backlight/
git clone https://github.com/WestleyK/rpi-backlight
cd rpi-backlight/pre-compiled/
sudo ./make.sh install
```

<br>
<br>

### Update:

Update the repo:
```
cd ~/raspberrypi-backlight/rpi-backlight/
make update
make
sudo make install
```

Or just update the binary:

```
curl https://raw.githubusercontent.com/WestleyK/rpi-backlight/master/easy-install.sh | sh
sudo mv -f rpi-backlight /usr/local/bin/
```
Basically the same commands as installing the binary.


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

<br>
<br>

## Usage:

```
pi@raspberrypi:~ $ rpi-backlight -help
Usage: rpi-backlight [OPTION]
      -help | --help (print help menu)
      [15-255] (adjust from: 15 to: 255)
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



