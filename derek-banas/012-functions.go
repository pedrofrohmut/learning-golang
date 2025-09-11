package main

import (
    "fmt"
)

func sayHello(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

func sum(x int, y int) int {
    return x + y
}

func sumAll(args ...int) int {
    var sum int = 0
    for _, n := range args {
        sum += n
    }
    return sum
}

func divide(x float64, y float64) (float64, error) {
    if y == 0 {
        return 0, fmt.Errorf("Cannot divide by zero")
    }
    return x / y, nil
}

func getArraySum(numbers []int) int {
    var sum int = 0
    for _, number := range numbers {
        sum += number
    }
    return sum
}

func F012() {
    sayHello("Bob")

    fmt.Println("My sum:", sum(34, 35))

    _, _ = divide(10.0, 3.0) // Normal

    var val, err = divide(10.0, 0.0) // Error
    if err != nil {
        fmt.Printf("ERROR: Error occured while dividing: %s\n", err)
    }
    fmt.Printf("Division: %f\n", val)

    // Variadic functions
    fmt.Printf("My other sum: %d\n", sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 3))

    fmt.Printf("Sum array: %d\n", getArraySum([]int { 35, 23, 11 }))
}
