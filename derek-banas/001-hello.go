package main

import (
    "fmt"
    "log"
    "bufio"
    "os"
)

/*
* Multile comment
 */

// Single line comment

func Fn001() {
    fmt.Println("What is your name?")
    var reader = bufio.NewReader(os.Stdin)
    var name, err = reader.ReadString('\n')
    if err != nil {
        log.Fatalf("ERROR: Error occured trying to read input from stdin: %s", err)
    }
    fmt.Println("Hello, ", name)
}
