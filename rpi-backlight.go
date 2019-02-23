// Created by: WestleyK
// email: westleyk@nym.hush.com
// Date: Feb 22, 2019
// https://github.com/WestleyK/rpi-backlight
// Version-1.1.4
//
// Designed and tested for raspberry pi with official 7 inch touchscreen. 
//
// The Clear BSD License
//
// Copyright (c) 2018-2019 WestleyK
// All rights reserved.
//
// This software is licensed under a Clear BSD License.
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

const (
    SCRIPT_VERSION string = "v-1.1.4"
    SCRIPT_DATE string = "Date: Feb 22, 2019"

    MIN_BRIGHTNESS string = "15"
    MAX_BRIGHTNESS string = "255"
    DEFAULT_ON string = "180"
    ADJUST_UP string = "25"
    ADJUST_DOWN string = "25"
    BRIGHTNESS_FILE string = "/sys/class/backlight/rpi_backlight/brightness"
    POWER_FILE string = "/sys/class/backlight/rpi_backlight/bl_power"

    colorReset string = "\x1b[0m"
    red string = "\x1b[31m"
//    yello string = "\x1b[33m"
//    blue string = "\x1b[34m"
//    pink string = "\x1b[35m"
//    white string = "\x1b[37m"
//    teal string = "\x1b[36m"
)

var (
    BRIGHT string = ""
)

func help_menu() {
    fmt.Printf("Usage: rpi-backlight [OPTION]\n")
    fmt.Printf("      --help         : print help menu.\n")
    fmt.Printf("      [%v-%v]       : adjust from: %v to: %s.\n", MIN_BRIGHTNESS, MAX_BRIGHTNESS, MIN_BRIGHTNESS, MAX_BRIGHTNESS)
    fmt.Printf("      -s, --sleep    : enter sleep mode, press <ENTER> to exit this mode.\n")
    fmt.Printf("      -u, --up       : adjust brightness up by: %v/%v.\n", ADJUST_UP, MAX_BRIGHTNESS)
    fmt.Printf("      -d, --down     : adjust brightness down by: %v/%v.\n", ADJUST_DOWN, MAX_BRIGHTNESS)
    fmt.Printf("      -c, --current  : print current brightness.\n")
    fmt.Printf("      -n, --on       : turn backlight on to: %v.\n", DEFAULT_ON)
    fmt.Printf("      --version      : print script version.\n")
    fmt.Printf("\n")
    fmt.Printf("Copyright (c) 2018-2019 WestleyK, All rights reserved.\n")
    fmt.Printf("This software is licensed under a Clear BSD License.\n")
    fmt.Printf("Source code: https://github.com/WestleyK/rpi-backlight\n")
    os.Exit(0)
}

func script_version() {
    fmt.Printf("%s, %s\n", SCRIPT_VERSION, SCRIPT_DATE)
    os.Exit(0)
}

func is_bright_file() {
    if _, err := os.Stat(BRIGHTNESS_FILE); os.IsNotExist(err) {
        fmt.Print(red, "ERROR: ", colorReset)
        fmt.Print("File does not exist:\n", BRIGHTNESS_FILE, "\n")
        os.Exit(1)
    }
}

func is_write() {
    file, err := os.Create(BRIGHTNESS_FILE)
    if os.IsPermission(err) {
        fmt.Print(red, "ERROR: ", colorReset)
        fmt.Print("Unable to write to: ", POWER_FILE, "\n")
        fmt.Println(err)
        os.Exit(1)
    }
    if err != nil {
        fmt.Print(red, "ERROR: ", colorReset)
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()
}

func is_power_file() {
    if _, err := os.Stat(POWER_FILE); os.IsNotExist(err) {
        fmt.Print(red, "ERROR: ", colorReset)
        fmt.Print("File does not exist:\n", POWER_FILE, "\n")
        os.Exit(1)
    }
    file, err := os.Create(POWER_FILE)
    if err != nil {
        if os.IsPermission(err) {
            fmt.Print(red, "ERROR: ", colorReset)
            fmt.Print("Unable to write to: ", POWER_FILE, "\n")
            fmt.Println(err)
            os.Exit(1)
        }
        fmt.Print(red, "ERROR: ", colorReset)
        fmt.Println(err)
        os.Exit(1)



    }
    defer file.Close()
}

func is_power_file_perm() {
    file_perm, err := os.Create(POWER_FILE)
    if err != nil {
        if os.IsPermission(err) {
            fmt.Print(red, "ERROR: ", colorReset)
            fmt.Print("Unable to write to: ", POWER_FILE, "\n")
            fmt.Println(err)
            os.Exit(1)
        }
        fmt.Print(red, "ERROR: ", colorReset)
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
        fmt.Print(red, "ERROR: ", colorReset)
        fmt.Println(err)
        os.Exit(1)
    }
    defer file_off.Close()
    // write 1 to POWER_FILE
    fmt.Fprintf(file_off, "1")

    // wait until user presses <ENTER>
    reader := bufio.NewReader(os.Stdin)
    _, err = reader.ReadString('\n')
    if err != nil {
        fmt.Print(red, "ERROR: ", colorReset)
        fmt.Println(err)
        os.Exit(1)
    }

    // then reopen the file, and write 0 to POWER_FILE
    file_on, err := os.Create(POWER_FILE)
    if err != nil {
        fmt.Print(red, "ERROR: ", colorReset)
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
        fmt.Print(red, "ERROR: ", colorReset)
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
        fmt.Print(red, "ERROR: ", colorReset)
        fmt.Println(err)
        os.Exit(1)
    }
    defer file_bright.Close()
    fmt.Fprintf(file_bright, DEFAULT_ON)

    file_power, err := os.Create(POWER_FILE)
    if err != nil {
        fmt.Print(red, "ERROR: ", colorReset)
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


        if OPTION == "--help" {
            help_menu()
        } else if OPTION == "-u" || OPTION == "--up" {
            adjust_up()
        } else if OPTION == "-d" || OPTION == "--down" {
            adjust_down()
        } else if OPTION == "-s" || OPTION == "--sleep" {
            sleep_mode()
        } else if OPTION == "-c" || OPTION == "--current" {
            current_bright()
        } else if OPTION == "-n" || OPTION == "--on" {
            turn_on()
        } else if OPTION == "--version" {
            script_version()
        } else if _, err := strconv.Atoi(OPTION); err == nil {
            BRIGHT = OPTION
            adjust_bright()
        } else {
            fmt.Print("Option not found!  :P  ", OPTION, "\n")
            fmt.Print("Try:  $ rpi-backlight --help  (for help)\n")
            os.Exit(1)
        }
    }
    current_bright()

}

//
// End source code
//
