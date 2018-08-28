#!/bin/sh
# Created by: Westley K
# email: westley@sylabs.io
# Date: Aug 27, 2018
# version-1.0.2
# https://github.com/WestleyK/rpi-backlight
#
# MIT License
#

set -e

SCRIPT_URL="https://github.com/WestleyK/rpi-backlight/blob/master/pre-compiled/raspberry-pi/rpi-backlight?raw=true"
SCRIPT_NAME="rpi-backlight"
INSTALL_PATH="/usr/local/bin/"

echo "Install script version:"
echo "version-1.0.0"

wget $SCRIPT_URL
mv rpi-backlight?raw=true $SCRIPT_NAME
chmod +x $SCRIPT_NAME
mv $SCRIPT_NAME $INSTALL_PATH

#
# End install script
#
