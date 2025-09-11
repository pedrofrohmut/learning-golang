package main

import (
    "fmt"
)

func Fn009() {

    // Golang accepts i++, i-- but not ++i, --i

    for i := 0; i < 10; i++ { // The variable i is local to the loop
        fmt.Println(i)
    }

    for i := 10; i >= 0; i-- {
        fmt.Println(i)
    } // Boom!

    // There is no while in golang, so you use for loop too.

    var age = 10
    for age < 65 {
        fmt.Printf("Still working at the age of %d\n", age)
        age++
    }
    fmt.Println("Can retire now finally!")

    // The do while in golang is like a hack: do a endless loop with a if/break at the bottom

    var count = 0
    for {
        fmt.Println("Count:", count)
        count++
        if (count > 5) { break } // This is the golang do-while loop
    }

    // Ranges
    for i := range 10 { // Goes from 0 to 9
        fmt.Println(i)
    }

    var arr1 = []int { 1, 2, 3, 4, 5}
    for i, number := range arr1 {
        fmt.Printf("[%d]: %d\n", i, number)
    }
}
