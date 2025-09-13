package main

import (
    "fmt"
    "os"
    "strconv"
)

func Fn015() {
    fmt.Println(os.Args)
    var args = os.Args[1:]
    var iargs = []int {}
    for _, arg := range args {
        var value, err = strconv.Atoi(arg)
        if err != nil {
            // log.Fatalf("Error to covert args to integer: %s", err)
            panic(err)
        }
        iargs = append(iargs, value)
    }

    var max = iargs[0]
    var min = iargs[0]
    for _, n := range iargs[1:] {
        if n > max {
            max = n
        }
        if n < min {
            min = n
        }
    }
    fmt.Printf("Max : %d, Min: %d", max, min)
}
