package main

import (
    "fmt"
    "os"
    "strconv"
    "log"
    "bufio"
    "errors"
)

func Fn014() {
    // Create a file
    var f, err = os.Create("data.txt")
    if err != nil {
        log.Fatalf("ERROR: Error to create a file: %s", err)
    }
    defer f.Close()

    var primes = []int { 2, 3, 5, 7, 11, 13 }
    var strPrimes = []string {}

    for _, prime := range primes {
        strPrimes = append(strPrimes, strconv.Itoa(prime))
    }

    // Write primes to file
    for _, num := range strPrimes {
        var _, err = f.WriteString(num + "\n")
        if err != nil {
            log.Fatalf("Error to write to file: %s", err)
        }
    }

    // Open the file
    var f2, errf2 = os.Open("data.txt")
    if errf2 != nil {
        log.Fatalf("ERROR: Error to open the file: %s", errf2)
    }
    defer f2.Close()

    // Read from opened file
    var scanner = bufio.NewScanner(f2)
    for scanner.Scan() {
        fmt.Printf("Prime: %s\n", scanner.Text())
    }

    // Checks scanner for error
    if err := scanner.Err(); err != nil {
        log.Fatalf("ERROR: Error to read from file %s", err)
    } else {
        log.Print("No errors on the scanner")
    }

    /*
       // Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
       O_RDONLY int = syscall.O_RDONLY // open the file read-only.
       O_WRONLY int = syscall.O_WRONLY // open the file write-only.
       O_RDWR   int = syscall.O_RDWR   // open the file read-write.
       // The remaining values may be or'ed in to control behavior.
       O_APPEND int = syscall.O_APPEND // append data to the file when writing.
       O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
       O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
       O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
       O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
    */

    // Check if file exist
    var _, errStat = os.Stat("data.txt")
    if errors.Is(errStat, os.ErrNotExist) {
        fmt.Println("File doesn't exist")
        return
    }

    // Open file
    var f3, errOpen = os.OpenFile("data.txt", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
    if errOpen != nil {
        log.Fatalf("ERROR: Could not open the file: %s", errOpen)
    }
    defer f3.Close()

    // Write to file
    if _, errWrite := f3.WriteString("17\n"); errWrite != nil {
        log.Fatal(errWrite)
    }
}
