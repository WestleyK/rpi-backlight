# Created by: WestleyK
# email: westley@sylabs.io
# Date: Feb 22, 2019
# https://github.com/WestleyK/rpi-backlight
# Version-1.0.0
#
# Designed and tested for raspberry pi with official 7 inch touchscreen. 
#
# The Clear BSD License
#
# Copyright (c) 2018-2019 WestleyK
# All rights reserved.
#
# This software is licensed under a Clear BSD License.
#

TARGET = rpi-backlight
PREFIX = /usr/local/bin

MAIN = rpi-backlight.go

.PHONY:
all: $(TARGET)

.PHONY:
$(TARGET): $(MAIN)
	go build

.PHONY:
install: $(TARGET)
	chmod +x $(TARGET)
	cp -f $(TARGET) $(PREFIX)

.PHONY:
uninstall: $(PREFIX)/$(TARGET)
	rm -f $(PREFIX)/$(TARGET)

#
# End Makefile
#
