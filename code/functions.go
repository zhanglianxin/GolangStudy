package main

import "fmt"

func main() {
    var a int = 1
    var b string = "a"

    c, d := swap(a, b)
    fmt.Println(c, d)

    fmt.Println()

    var x, y int = 100, 200

    fmt.Printf("Call by value: ")
    swapcallbyvalue(x, y)
    fmt.Printf("x is %d, y is %d", x, y)

    fmt.Println()

    fmt.Printf("Call by reference: ")
    swapcallbyreference(&x, &y)
    fmt.Printf("x is %d, y is %d", x, y)
}

func swap(x int, y string) (string, int) {
    return y, x
}

// Call by value
func swapcallbyvalue(x, y int) {
    var temp int;

    temp = x
    x = y
    y = temp
}

// Call by reference
func swapcallbyreference(x *int, y *int) {
    var temp int;

    temp = *x
    *x = *y
    *y = temp
}
