#!/bin/sh
# Created by: Westley K
# email: westley@sylabs.io
# Date: Aug 29, 2018
# version-1.0.6
# https://github.com/WestleyK/rpi-backlight
#
# MIT License
#

set -e

SCRIPT_URL="https://github.com/WestleyK/rpi-backlight/blob/master/pre-compiled/raspberry-pi/rpi-backlight?raw=true"
SCRIPT_NAME="rpi-backlight"
INSTALL_PATH="/usr/local/bin/"
KERNEL_VERSION="4.14.50-v7+"

echo "Install script version: version-1.0.6"
echo 

KERNEL=` uname -r `
if [ "$KERNEL" != "$KERNEL_VERSION" ]; then
    printf "\033[0;31mFATAL ERROR: \033[0m"
    echo "Kernel not supported!"
    echo "Supported kernel version: $KERNEL_VERSION"
    echo "Your kernel version: $KERNEL"
    echo 
    echo "!!FAIL: Install failed!"
    exit 1
fi

echo "Downloading file..."
wget -O $SCRIPT_NAME $SCRIPT_URL >/dev/null 2>&1
chmod +x $SCRIPT_NAME

echo "!!SUCCESS: install successful."
echo 
echo "########## rpi-backlight installed to the current directory  ##########"
echo
echo ">> To finish the install, do:"
echo " $ sudo mv rpi-backlight /usr/local/bin/"

#
# End install script
#
