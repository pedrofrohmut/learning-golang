package main

import (
    "fmt"
)

func doubleArray(arr *[4]int) {
    for i := range *arr {
        arr[i] *= 2
    }
}

func changeVal(n int) {
    n += 10
}

func changeVal2(n *int) {
    *n += 10
}

func getArrayAvarage(nums ...float64) float64 {
    var sum float64 = 0.0
    for _, num := range nums {
        sum += num
    }
    return sum / float64(len(nums))
}

func main() {
    var n int = 10

    fmt.Printf("Value of n before changeVal: %d\n", n)

    changeVal(n) // The variable is passed by value here so n wont change

    fmt.Printf("Value of n after changeVal: %d\n", n)

    changeVal2(&n) // The variable is passed by reference so the value changes

    fmt.Printf("Value of n after changeVal2: %d\n", n)

    var n2 = &n

    fmt.Printf("Before increment n2 => Value of n: %d, n2: %d\n", n, *n2)

    *n2 += 10 // Since n2 is a pointer both the values will change

    fmt.Printf("After increment n2 => Value of n: %d, n2: %d\n", n, *n2)

    // Pointer to array
    var parr = [4]int { 1, 2, 3, 4 }

    doubleArray(&parr)

    fmt.Println(parr)

    var islice = []float64 { 11, 13, 17 }
    var avg = getArrayAvarage(islice...)
    fmt.Printf("Avarage: %.3f\n", avg)
}
