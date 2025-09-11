package main

import (
    "fmt"
    "reflect"
    "strconv"
    "log"
)

func fn003() {
    // Data types
    fmt.Println(reflect.TypeOf(25))
    fmt.Println(reflect.TypeOf(3.14))
    fmt.Println(reflect.TypeOf(true))
    fmt.Println(reflect.TypeOf("Hello"))
    fmt.Println(reflect.TypeOf('F'))

    // Convert
    var cv1 = 1.5
    var cv2 = int(cv1) // Floors the number, removes the decimal part
    fmt.Println(cv2)
    var cv3 = "5000"

    var cv4, errAtoi = strconv.Atoi(cv3) // Atoi - ASCII to Integer
    if errAtoi != nil {
        log.Fatalf("ERROR: Could not convert string to integer: %s", errAtoi)
    }
    fmt.Println(cv3, reflect.TypeOf(cv3), cv4, reflect.TypeOf(cv4))

    var cv5 = 5000
    var cv6 = strconv.Itoa(cv5) // Atoi - ASCII to Integer
    fmt.Println(cv5, reflect.TypeOf(cv5), cv6, reflect.TypeOf(cv6))

    var cv7 = "3.14"
    if cv8, err := strconv.ParseFloat(cv7, 10); err != nil {
        fmt.Println(cv8, reflect.TypeOf(cv8))
    }

    var cv9 = fmt.Sprintf("%f", 3.14)
    fmt.Println(cv9, reflect.TypeOf(cv9))

    // %d - Integer with leading zeros %04d, leading spaces %4d
    // %c - Char
    // %f - Float with leading zeros %04f, 2 after the point %.2f or two together %04.2f, no decimals %.f
    // %t - Boolean
    // %s - String
    // %o - Base 8
    // %x - Base 16
    // %v - Guesses based on data type
    // %T - Type of supplied value
}
