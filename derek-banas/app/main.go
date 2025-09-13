package main

import (
    "fmt"
    "reflect"
    "strings"
    "log"
    "example/project/mypackage"
)

func main() {
    fmt.Println("Hello,", stuff.Name)
    var primes = []int { 2, 3, 5, 7, 9, 11, 13 }
    var strPrimes = stuff.IntArrToStrArr(primes)
    for _, x := range strPrimes {
        fmt.Println(x)
    }
    fmt.Println(strPrimes)
    fmt.Println(reflect.TypeOf(strPrimes))
    fmt.Println(strings.Join(strPrimes, ":"))

    // Encapsulation - Date fields are not visible here
    var date = stuff.Date {}
    var errMonth = date.SetMonth(9)
    if errMonth != nil {
        log.Fatal(errMonth)
    }
    var errDay = date.SetDay(13)
    if errDay != nil {
        log.Fatal(errDay)
    }
    var errYear = date.SetYear(2025)
    if errYear != nil {
        log.Fatal(errYear)
    }

    // You cannot access the day, month, year here. They are not visible
    // fmt.Printf("Date is %d/%d/%d", date.day, date.month, date.year) // Error here
    fmt.Printf("Date is %02d/%02d/%04d", date.Day(), date.Month(), date.Year())
}
