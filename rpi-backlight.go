// Created by: Westley K
// email: westley@sylabs.io
// Date: Aug 28, 2018
// https://github.com/WestleyK/rpi-backlight
// Version-1.1.2
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
    "strings"
    "io/ioutil"
)

var SCRIPT_VERSION = "version-1.1.2"
var SCRIPT_DATE = "Date: Aug 28, 2018"

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
    fmt.Print("      [", MIN_BRIGHTNESS, "-", MAX_BRIGHTNESS, "] (adjust from: ", MIN_BRIGHTNESS, " to: ", MAX_BRIGHTNESS, ")\n")
    fmt.Print("      -s | -sleep (enter sleep mode, press <ENTER> to exit this mode)\n")
    fmt.Print("      -u | -up (adjust brightness up by: ", ADJUST_UP, "/", MAX_BRIGHTNESS, ")\n")
    fmt.Print("      -d | -down (adjust brightness down by: ", ADJUST_DOWN, "/", MAX_BRIGHTNESS, ")\n")
    fmt.Print("      -c | -current (print current brightness)\n")
    fmt.Print("      -n | -on (turn backlight on to: ", DEFAULT_ON, "\n")
    fmt.Print("      -i | -info (print info)\n")
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

func is_bright_file() {
    if _, err := os.Stat(BRIGHTNESS_FILE); os.IsNotExist(err) {
        fmt.Print("\033[0;31mERROR: \033[0m")
        fmt.Print("File does not exist:\n", BRIGHTNESS_FILE, "\n")
        os.Exit(1)
    }
}

func is_write() {
    file, err := os.Create(BRIGHTNESS_FILE)
    if err != nil {
        if os.IsPermission(err) {
            fmt.Print("\033[0;31mERROR: \033[0m")
            fmt.Print("Unable to write to: ", BRIGHTNESS_FILE, "\n")
            fmt.Println(err)
            os.Exit(1)
        }
        fmt.Print("\033[0;31mERROR: \033[0m")
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()
}

func is_power_file() {
    if _, err := os.Stat(POWER_FILE); os.IsNotExist(err) {
        fmt.Print("\033[0;31mERROR: \033[0m")
        fmt.Print("File does not exist:\n", POWER_FILE, "\n")
        os.Exit(1)
    }
    file, err := os.Create(POWER_FILE)
    if err != nil {
        if os.IsPermission(err) {
            fmt.Print("\033[0;31mERROR: \033[0m")
            fmt.Print("Unable to write to: ", POWER_FILE, "\n")
            fmt.Println(err)
            os.Exit(1)
        }
        fmt.Print("\033[0;31mERROR: \033[0m")
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()
}

func is_power_file_perm() {
    file_perm, err := os.Create(POWER_FILE)
    if err != nil {
        if os.IsPermission(err) {
            fmt.Print("\033[0;31mERROR: \033[0m")
            fmt.Print("Unable to write to: ", POWER_FILE, "\n")
            fmt.Println(err)
            os.Exit(1)
        }
        fmt.Print("\033[0;31mERROR: \033[0m")
        fmt.Println(err)
        os.Exit(1)
    }
    defer file_perm.Close()
}

func sleep_mode() {
    is_power_file()
    is_power_file_perm()

    // sleep 1s, then write 1 to POWER_FILE
    fmt.Print("Press <ENTER> to exit this mode:\n")
    time.Sleep(1 * time.Second)
    file_off, err := os.Create(POWER_FILE)
    if err != nil {
        fmt.Print("\033[0;31mERROR: \033[0m")
        fmt.Println(err)
        os.Exit(1)
    }
    defer file_off.Close()
    // write 1 to POWER_FILE
    fmt.Fprintf(file_off, "1")

    // wait until user presses <ENTER>
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    // do nothing with the text
    _ = text

    // then reopen the file, and write 0 to POWER_FILE
    file_on, err := os.Create(POWER_FILE)
    if err != nil {
        fmt.Print("\033[0;31mERROR: \033[0m")
        fmt.Println(err)
        os.Exit(1)
    }
    defer file_on.Close()
    // write 0 to POWER_FILE
    fmt.Fprintf(file_on, "0")

    os.Exit(0)
}

func current_bright() {
    is_bright_file()

    b, err := ioutil.ReadFile(BRIGHTNESS_FILE)
    if err != nil {
        fmt.Print(err)
    }

    CURRENT_STRING := string(b)

    CURRENT_STRING = strings.TrimSuffix(CURRENT_STRING, "\n")
    CURRENT, err := strconv.Atoi(CURRENT_STRING)

    fmt.Print("Current brightness: ", CURRENT, "\n")
    os.Exit(0)
}

func write_file() {
    is_bright_file()
    is_write()

    file, err := os.Create(BRIGHTNESS_FILE)
    if err != nil {
        fmt.Print("\033[0;31mERROR: \033[0m")
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()
    fmt.Fprintf(file, BRIGHT)
    fmt.Print("Current brightness: ", BRIGHT, "\n")
    os.Exit(0)
}

func adjust_up() {
    is_bright_file()

    b, err := ioutil.ReadFile(BRIGHTNESS_FILE)
    if err != nil {
        fmt.Print(err, "\n")
        os.Exit(1)
    }

    CURRENT_STRING := string(b)

    CURRENT_STRING = strings.TrimSuffix(CURRENT_STRING, "\n")
    CURRENT, _ := strconv.Atoi(CURRENT_STRING)
    ADJUST_UP, _ := strconv.Atoi(ADJUST_UP)
    MAX_BRIGHTNESS, _ := strconv.Atoi(MAX_BRIGHTNESS)

    CURRENT += ADJUST_UP
    if CURRENT >= MAX_BRIGHTNESS {
        fmt.Print("Max Brightness!\n")
        CURRENT = MAX_BRIGHTNESS
    }
    BRIGHT = strconv.Itoa(CURRENT)
    write_file()
}

func adjust_down() {
    is_bright_file()

    b, err := ioutil.ReadFile(BRIGHTNESS_FILE)
    if err != nil {
        fmt.Print(err, "\n")
        os.Exit(1)
    }

    CURRENT_STRING := string(b)

    CURRENT_STRING = strings.TrimSuffix(CURRENT_STRING, "\n")
    CURRENT, _ := strconv.Atoi(CURRENT_STRING)
    ADJUST_DOWN, _ := strconv.Atoi(ADJUST_DOWN)
    MIN_BRIGHTNESS, _ := strconv.Atoi(MIN_BRIGHTNESS)

    CURRENT -= ADJUST_DOWN
    if CURRENT <= MIN_BRIGHTNESS {
        fmt.Print("Min Brightness!\n")
        CURRENT = MIN_BRIGHTNESS
    }
    BRIGHT = strconv.Itoa(CURRENT)
    write_file()
}

func turn_on() {
    is_bright_file()
    is_write()
    is_power_file()

    file_bright, err := os.Create(BRIGHTNESS_FILE)
    if err != nil {
        fmt.Print("\033[0;31mERROR: \033[0m")
        fmt.Println(err)
        os.Exit(1)
    }
    defer file_bright.Close()
    fmt.Fprintf(file_bright, DEFAULT_ON)

    file_power, err := os.Create(POWER_FILE)
    if err != nil {
        fmt.Print("\033[0;31mERROR: \033[0m")
        fmt.Println(err)
        os.Exit(1)
    }
    defer file_power.Close()
    fmt.Fprintf(file_power, "0")
    os.Exit(0)
}

func adjust_bright() {
    // check if theres a '-'
    if strings.Contains(BRIGHT, "-") {
        fmt.Print("Only whole numbers! ", BRIGHT, "\n")
        os.Exit(1)

    }

    // convert strings to int
    CURRENT, _ := strconv.Atoi(BRIGHT)
    MIN_BRIGHTNESS, _ := strconv.Atoi(MIN_BRIGHTNESS)
    MAX_BRIGHTNESS, _ := strconv.Atoi(MAX_BRIGHTNESS)

    // check if in range
    if CURRENT >= MIN_BRIGHTNESS && CURRENT <= MAX_BRIGHTNESS {
        BRIGHT = strconv.Itoa(CURRENT)
        write_file()
        os.Exit(0)
    }

    fmt.Print("Not a valid number: ", CURRENT, "\n")
    fmt.Print("Your options: [", MIN_BRIGHTNESS, "-", MAX_BRIGHTNESS, "]\n")
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
        } else if OPTION == "-c" || OPTION == "-current" {
            current_bright()
        } else if OPTION == "-n" || OPTION == "-on" {
            turn_on()
        } else if OPTION == "-i" || OPTION == "-info" {
            info_script()
        } else if OPTION == "-v" || OPTION == "-version" || OPTION == "--version" {
            script_version()
        } else if _, err := strconv.Atoi(OPTION); err == nil {
            BRIGHT = OPTION
            adjust_bright()
        } else {
            fmt.Print("Option not found!  :P  ", OPTION, "\n")
            fmt.Print("Try:  $ rpi-backlight -help  (for help)\n")
            os.Exit(1)
        }
    }
    current_bright()

}



//
// End source code
//


