package main

import (
    "time"
    "fmt"
)

func main() {
    var now = time.Now()

    var isodate = fmt.Sprintf("%d-%02d-%dT%d.%d.%d.%dZ",
        now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.UnixMilli() % 1000)

    fmt.Println(isodate)
}
