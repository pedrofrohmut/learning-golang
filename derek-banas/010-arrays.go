package main

import (
    "fmt"
)

func Fn010() {
    var arr = [5]int { 13, 42, 69, 144, 420 } // Init an int array of size 5 with values

    arr[1] = '?' // Change the value

    _ = arr[0] // Get the value at index 0

    _ = len(arr) // Get the length of the array


    for i := 0; i < len(arr); i++ { // Traditional for loop for reference
        fmt.Println(i)
    }

    for i, number := range arr { // Easier looping arrays with range
        fmt.Printf("[%d]: %d\n", i, number)
    }

    // Multidimensional arrays
    var matrix = [2][2]int {
        { 1, 2 },
        { 3, 4 },
    }
    for i := range matrix {
        for j := range matrix[i] {
            fmt.Printf("[%d][%d]: %d\n", i, j, matrix[i][j])
        }
    }

    // Strings
    for _, value := range []rune("abcdef") {
        fmt.Printf("Rune array: %d\n", value)
    }

    var bytearr = []byte { 'a', 'b', 'c', 'd', 'e', 'f' }
    _ = string(bytearr)
}
