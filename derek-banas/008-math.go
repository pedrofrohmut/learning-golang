package main

import (
    "math/rand"
    "math"
    "fmt"
)

func Fn008() {
    _ = 5 + 4
    _ = 5 - 4
    _ = 5 * 4
    _ = 5 / 4
    _ = 5 % 4 // Rest

    var i = 1
    i++
    i += 3
    i = i / 3

    var randNum = rand.Int()
    fmt.Println(randNum)

    _ = math.Abs(-10)
    _ = math.Pow(4, 2)
    _ = math.Sqrt(4)
    _ = math.Cbrt(4)
    _ = math.Ceil(4)
    _ = math.Floor(4)
    _ = math.Round(4)
    _ = math.Log2(4)
    _ = math.Log10(4)
    _ = math.Log(4)
    _ = math.Max(4, 8)
    _ = math.Min(4, 8)

    // Convert 90 degrees to radians
    var r90 = 90 * math.Pi / 180
    _ = r90 * (180 / math.Pi)
}
