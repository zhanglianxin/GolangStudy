package main

import "fmt"

func main() {
    var x interface{}

    switch i := x.(type) {
        case nil:
            fmt.Printf("type of x:%T\n", i)
        case int:
            fmt.Printf("x is int\n")
        case float64:
            fmt.Printf("x is float64\n")
        case func(int) float64:
            fmt.Printf("x is func(int)\n")
        case bool, string:
            fmt.Printf("x is bool or string\n")
        default:
            fmt.Printf("don't konw the type\n")
    }
}
