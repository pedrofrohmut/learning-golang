package main

import (
    "fmt"
)

// User defined types
type Tsp float64 // Tea Spoon
type TBs float64 // Table Spoon
type ML float64 // Milliliter

// When you use defined types you can use then Associative methods
func (this Tsp) ToMLs() ML { return ML(this * 5.92) }
func (this TBs) ToMLs() ML { return ML(this * 17.76) }

func teaSpoonToMilliliter(value Tsp) ML {
    return ML(value * 5.92)
}

func tableSpoonToMilliliter(value TBs) ML {
    return ML(value * 17.76)
}

func Fn019() {
    // With normal defined methods
    fmt.Printf("3 tea spoons is %.f ml\n", teaSpoonToMilliliter(3))
    fmt.Printf("3 table spoons is %.f ml\n", tableSpoonToMilliliter(3))

    // With associative methods
    fmt.Printf("3 tea spoons is %.f ml\n", Tsp(3).ToMLs())
    fmt.Printf("3 table spoons is %.f ml\n", TBs(3).ToMLs())
}
