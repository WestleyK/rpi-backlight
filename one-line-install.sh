#!/bin/sh
# Created by: Westley K
# email: westley@sylabs.io
# Date: Aug 27, 2018
# version-1.0.4
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
echo "version-1.0.4"

KERNEL=` uname -r `
if [ "$KERNEL" != "$KERNEL_VERSION" ]; then
    echo "Kernel not supported!"
    echo "Your kernel version: $KERNEL"
    echo "Supported kernel version: $KERNEL_VERSION"
    echo 
    echo "Install failed!"
    exit 1
fi

wget $SCRIPT_URL
mv rpi-backlight?raw=true $SCRIPT_NAME
chmod +x $SCRIPT_NAME
mv -f $SCRIPT_NAME $INSTALL_PATH

echo 
echo "########## rpi-backlight installed to /usr/local/bin/ ##########"
echo

#
# End install script
#
