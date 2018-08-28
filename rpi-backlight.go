// Created by: Westley K
// email: westley@sylabs.io
// Date: Aug 27, 2018
// https://github.com/WestleyK/rpi-backlight
// Version-1.0.2
//
// Designed and tested for raspberry pi with official 7 inch touchdcreen. 
//
//
// MIT License
//
// Copyright (c) 2018 WestleyK
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//


package main


import (
    "os"
    "fmt"
    "strconv"
    "time"
    "bufio"
    "io/ioutil"
)

var SCRIPT_VERSION = "version-1.0.2"
var SCRIPT_DATE = "Date: Aug 27, 2018"

var MIN_BRIGHTNESS = "15"
var MAX_BRIGHTNESS = "255"
var DEFAULT_ON = "180"
var ADJUST_UP = "25"
var ADJUST_DOWN = "25"
var BRIGHTNESS_FILE = "/sys/class/backlight/rpi_backlight/brightness"
var POWER_FILE = "/sys/class/backlight/rpi_backlight/bl_power"

var BRIGHT = ""

func help_menu() {
    fmt.Print("Usage: rpi-backlight [OPTION]\n")
    fmt.Print("      -help | --help (print help menu)\n")
    fmt.Print("      -u | -up (adjust brightness up by: ", ADJUST_UP, "/", MAX_BRIGHTNESS, ")\n")
    fmt.Print("      -d | -down (adjust brightness down by: ", ADJUST_DOWN, "/", MAX_BRIGHTNESS, ")\n")
    fmt.Print("      -info | --info (print info)\n")
    fmt.Print("      -version | --version (print version)\n")
    fmt.Print("Source code: https://github.com/WestleyK/rpi-backlight\n")
    os.Exit(0)
}

func script_version() {
    fmt.Println(SCRIPT_VERSION)
    fmt.Println(SCRIPT_DATE)
    os.Exit(0)
}

func info_script() {
    info()
    os.Exit(0)
}

func sleep_mode() {
    fmt.Print("Press <ENTER> to exit this mode:\n")
    time.Sleep(1 * time.Second)
    errF := ioutil.WriteFile(POWER_FILE, []byte("1"), 0644)
    if errF != nil {
        fmt.Println(errF)
    }


    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    fmt.Println(text)

    errN := ioutil.WriteFile(POWER_FILE, []byte("0"), 0644)
    if errN != nil {
        fmt.Println(errN)
    }
    os.Exit(0)
}

func write_file() {
    if _, err := os.Stat(BRIGHTNESS_FILE); os.IsNotExist(err) {
        fmt.Print("File does not exist:\n", BRIGHTNESS_FILE)
        os.Exit(1)
    }
    err := ioutil.WriteFile(BRIGHTNESS_FILE, []byte(BRIGHT), 0644)
    if err != nil {
        if os.IsPermission(err) {
            fmt.Println("Unable to write to ", BRIGHTNESS_FILE)
            fmt.Println(err)
            os.Exit(1)
        }
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Print("Current brightness: ", BRIGHT, "\n")
    os.Exit(0)
}

func adjust_up() {
    b, err := ioutil.ReadFile(BRIGHTNESS_FILE)
    if err != nil {
        fmt.Print(err)
    }

    CURRENT_STRING := string(b)

    CURRENT, err := strconv.Atoi(CURRENT_STRING)
    ADJUST_UP, err := strconv.Atoi(ADJUST_UP)
    MAX_BRIGHTNESS, err := strconv.Atoi(MAX_BRIGHTNESS)

    CURRENT += ADJUST_UP
    if CURRENT >= MAX_BRIGHTNESS {
        fmt.Print("Max Brightness!\n")
        CURRENT = MAX_BRIGHTNESS
    }
    BRIGHT = strconv.Itoa(CURRENT)
    write_file()
}

func adjust_down() {
    b, err := ioutil.ReadFile(BRIGHTNESS_FILE)
    if err != nil {
        fmt.Print(err)
    }

    CURRENT_STRING := string(b)

    CURRENT, err := strconv.Atoi(CURRENT_STRING)
    ADJUST_DOWN, err := strconv.Atoi(ADJUST_DOWN)
    MIN_BRIGHTNESS, err := strconv.Atoi(MIN_BRIGHTNESS)

    CURRENT -= ADJUST_DOWN
    if CURRENT <= MIN_BRIGHTNESS {
        fmt.Print("Min Brightness!\n")
        CURRENT = MIN_BRIGHTNESS
    }
    BRIGHT = strconv.Itoa(CURRENT)
    write_file()
}

func turn_on() {
    errB := ioutil.WriteFile(BRIGHTNESS_FILE, []byte(DEFAULT_ON), 0644)
    if errB != nil {
        fmt.Println(errB)
    }
    errP := ioutil.WriteFile(POWER_FILE, []byte("0"), 0644)
    if errP != nil {
        fmt.Println(errP)
    }
    os.Exit(0)
}

func main() {

    if len(os.Args[1:]) >= 1 {

        if len(os.Args[1:]) >= 2 {
            fmt.Print("Only one argument!\n")
            os.Exit(1)
        }

        OPTION := os.Args[1]


        if OPTION == "-h" || OPTION == "-help" || OPTION == "--help" {
            help_menu()
        } else if OPTION == "-u" || OPTION == "-up" {
            adjust_up()
        } else if OPTION == "-d" || OPTION == "-down" {
            adjust_down()
        } else if OPTION == "-s" || OPTION == "-sleep" {
            sleep_mode()
        } else if OPTION == "-n" || OPTION == "-on" {
            turn_on()
        } else if OPTION == "-i" || OPTION == "-info" {
            info_script()
        } else if OPTION == "-v" || OPTION == "-version" || OPTION == "--version" {
            script_version()
        } else {
            fmt.Print("Option not found!  :P  ", OPTION, "\n")
            fmt.Print("Try:  $ rpi-backlight -help  (for help)\n")
            os.Exit(1)
        }
    }

}



//
// End source code
//


