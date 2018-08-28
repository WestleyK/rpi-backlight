#!/bin/sh
# Created by: Westley K
# email: westley@sylabs.io
# Date: Aug 28, 2018
# version-1.0.5
# https://github.com/WestleyK/rpi-backlight
#
# MIT License
#

set -e

SCRIPT_URL="https://github.com/WestleyK/rpi-backlight/blob/master/pre-compiled/raspberry-pi/rpi-backlight?raw=true"
SCRIPT_NAME="rpi-backlight"
INSTALL_PATH="/usr/local/bin/"

KERNEL_VERSION="4.14.50-v7+"

echo "Install script version:"
echo "version-1.0.5"
echo 

KERNEL=` uname -r `
if [ "$KERNEL" != "$KERNEL_VERSION" ]; then
    printf "\033[0;31mFATAL ERROR: \033[0m"
    echo "Kernel not supported!"
    echo "Supported kernel version: $KERNEL_VERSION"
    echo "Your kernel version: $KERNEL"
    echo 
    echo "Install failed!"
    exit 1
fi

echo "Downloading file..."
wget $SCRIPT_URL &> /dev/null
mv rpi-backlight?raw=true $SCRIPT_NAME
chmod +x $SCRIPT_NAME

echo 
echo "########## rpi-backlight installed to /usr/local/bin/ ##########"
echo
echo ">> To finish the install, do:"
echo " $ sudo mv rpi-backlight /usr/local/bin/"

#
# End install script
#
