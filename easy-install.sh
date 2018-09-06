#!/bin/sh
# Created by: Westley K
# email: westley@sylabs.io
# Date: Aug 31, 2018
# version-1.0.9
# https://github.com/WestleyK/rpi-backlight
#
# MIT License
#

set -e

SCRIPT_URL="https://github.com/WestleyK/rpi-backlight/raw/master/pre-compiled/raspberry-pi/rpi-backlight"
SCRIPT_NAME="rpi-backlight"
INSTALL_PATH="/usr/local/bin/"
ARCH_VERSION="armv7l"

echo "Install script version: version-1.0.9"
echo 

ARCH_ON=` uname -m `
if [ "$ARCH_ON" != "$ARCH_VERSION" ]; then
    printf "\033[0;31mFATAL ERROR: \033[0m"
    echo "Architecture not supported!"
    echo "Supported architecture version: $ARCH_VERSION"
    echo "Your architecture version: $ARCH_ON"
    echo 
    echo "FAIL: Install failed!"
    exit 1
fi

echo "Downloading file..."
wget -O $SCRIPT_NAME $SCRIPT_URL
chmod +x $SCRIPT_NAME

echo 
echo
echo
echo "SUCCESS: install successful."
echo 
echo "########## $SCRIPT_NAME installed to the current directory  ##########"
echo
echo ">> To finish the install, do:"
echo " $ sudo mv rpi-backlight /usr/local/bin/"
exit 0

#
# End install script
#
