package main

import (
    "fmt"
    "strings"
)

func main() {
    var greeting = "Hello world!"

    fmt.Printf("String Length is: %d\n", len(greeting))

    greetings := []string{"Hello", "world!"}
    fmt.Println(strings.Join(greetings, " "))
}
