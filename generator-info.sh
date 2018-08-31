#!/bin/sh
#
# Created by: Westley K
# email: westley@sylabs.io
# Date: Aug 30, 2018
# https://github.com/WestleyK/rpi-backlight
# Version-1.0.2
#
# Designed and tested for raspberry pi with official 7 inch touchdcreen. 
#
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

FILE_INFO="rpi_backlight_info.go"
DATE=` date `
WHO=` whoami `
ON=` hostname `
WHERE=` pwd `
KERN=` uname -r `
ARCH=` uname -m `

touch $FILE_INFO
cat /dev/null > $FILE_INFO

cat << END_OF_FILE > $FILE_INFO

package main

import "fmt"

func info() {
    fmt.Print("Compiled date: $DATE\n")
    fmt.Print("Compiled by: $WHO\n")
    fmt.Print("Compiled on: $ON\n")
    fmt.Print("Compiled in: $WHERE\n")
    fmt.Print("Compiled on kernel: $KERN\n")
    fmt.Print("Compiled on architecture: $ARCH\n")

}

END_OF_FILE

#
# End script
#
