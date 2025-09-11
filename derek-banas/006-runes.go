package main

import (
    "fmt"
    "unicode/utf8"
)

func Fn006() {
    var rstr = "abcdef"
    fmt.Println("Rune count: ", utf8.RuneCountInString(rstr))
    for i, val := range rstr {
        fmt.Printf("%d : %#U : %c\n", i, val, val)
    }
}
