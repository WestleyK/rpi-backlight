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


### End README


<br>
<br>



