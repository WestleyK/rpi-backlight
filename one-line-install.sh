#!/bin/sh
# Created by: Westley K
# email: westley@sylabs.io
# Date: Aug 27, 2018
# version-1.0.1
# https://github.com/WestleyK/rpi-backlight
#
# MIT License
#
# Copyright (c) 2018 WestleyK
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.
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
