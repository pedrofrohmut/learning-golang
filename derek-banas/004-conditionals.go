package main

import (
    "fmt"
)

func fn004() {
    // Conditional Operators > < >= <= == !=
    // Logical Operators && || !
    var age = 8
    if age >= 1 && age <= 18 {
        fmt.Println("Important Birthday")
    } else if age == 21 || age == 50 {
        fmt.Println("Important Birthday")
    } else if age >= 65 {
        fmt.Println("Important Birthday")
    } else {
        fmt.Println("Not an Important Birthday")
    }
}
