package main

import (
    "fmt"
    "reflect"
)

func main() {
    var arr1 = make([]string, 5)
    arr1[0] = "Society"
    arr1[1] = "of"
    arr1[2] = "the"
    arr1[3] = "Simulated"
    arr1[4] = "Universe"

    for i, val := range arr1 {
        fmt.Printf("[%d]: %s\n", i, val)
    }

    var arr2 = [5]int { 13, 24, 42, 69, 144 }

    _ = arr2[0:2] // Gets the first 2 values
    _ = arr2[:3] // Gets the first 3 values (Can omit the left side if is 0)
    _ = arr2[1:] // Gets all values but the first one

    fmt.Println("First 2:")
    for _, val := range arr2[0:2] { // Gets first 2 elements
        fmt.Println(val)
    }

    fmt.Println("Last 3:")
    for _, val := range arr2[len(arr2) - 3:] { // Get last 3 elements
        fmt.Println(val)
    }

    var slice1 = arr2[:]
    slice1[1] = 666 // In golang when you alter the slice you also alter the original array
    arr2[3] = 555
    for i := range arr2 {
        fmt.Printf("Arr2: %d, Slice1: %d\n", arr2[i], slice1[i])
    }

    // Append to slice - You can append to slice but not to the array (fixed size)
    slice1 = append(slice1, 1234)

    // Creates a new string slice
    var arr3 = make([]string, 6)

    // Creates a new string array
    var arr4 = [6]string {}

    fmt.Printf("Arr3: %T, %s, Arr4: %T, %s\n", arr3, reflect.TypeOf(arr3), arr4, reflect.TypeOf(arr4))
}
