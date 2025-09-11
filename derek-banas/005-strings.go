package main

import (
    "fmt"
    "strings"
)

func fn005() {
    // Strings
    var sv1 = "A Word"
    var replacer = strings.NewReplacer("A", "Another")
    var sv2 = replacer.Replace(sv1)
    fmt.Println("#", sv1, "#", sv2)
    _ = len(sv2)
    _ = strings.Contains(sv2, "Another")
    _ = strings.Index(sv2, "A")
    _ = strings.ReplaceAll(sv2, "o", "0")
    _ = strings.TrimSpace("    foo bar            ")
    _ = strings.Split("0-1-2-3-4-5-6", "-")
    _ = strings.ToLower("Foo Bar")
    _ = strings.ToUpper("foo bar")
    _ = strings.HasPrefix("tacocat", "taco")
    _ = strings.HasSuffix("tacocat", "cat")
}
