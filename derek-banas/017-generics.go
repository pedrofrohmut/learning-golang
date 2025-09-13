package main

import (
    "fmt"
)

type MyConstraint interface {
    int | float64
}

func genericSum[T MyConstraint](numbers ...T) T {
    var sum T
    for _, n := range numbers {
        sum += n
    }
    return sum
}

func numberSum[T MyConstraint](x T, y T) T {
    return x + y
}

func Fn017() {
    var sum = genericSum(1.3, 1.4, 1.5)
    fmt.Println(sum)

    var sum2 = genericSum(1, 2, 3)
    fmt.Println(sum2)

    // This sum works for both integers and floats
    fmt.Println(numberSum(1.5, 3.5))
    fmt.Println(numberSum(1, 3))
}
