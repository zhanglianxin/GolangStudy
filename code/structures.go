package main

import (
    "fmt"
)

type Person struct {
    name string
    age int
}

func main() {

    var p1, p2 Person

    p1.name = "haha"
    p1.age = 18
    p2.name = "hehe"
    p2.age = 20

    fmt.Printf("%T\n", p1)
    fmt.Printf("%+v\n", p1)
    fmt.Printf("%#v\n", p1)
}
