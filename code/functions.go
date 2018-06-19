package main

import (
    "fmt"
    "math"
)

func main() {
    var a int = 1
    var b string = "a"

    c, d := swap(a, b)
    fmt.Println(c, d)

    var x, y int = 100, 200

    fmt.Println("Call by value: ")
    swapcallbyvalue(x, y)
    fmt.Printf("x is %d, y is %d", x, y)

    fmt.Println()

    fmt.Println("Call by reference: ")
    swapcallbyreference(&x, &y)
    fmt.Printf("x is %d, y is %d", x, y)

    fmt.Println()
    fmt.Println("Function as value: ")
    getSquareRoot := func (x float64) float64 {
        return math.Sqrt(x)
    }
    fmt.Println(getSquareRoot(9))

    fmt.Println("Function closure: ")
    nextNumber := getSequence()
    fmt.Println(nextNumber())
    fmt.Println(nextNumber())
    fmt.Println(nextNumber())

    nextNumber1 := getSequence()
    fmt.Println(nextNumber1())
    fmt.Println(nextNumber1())

    fmt.Println("Methods: (special types of functions)")
    circle := Circle{x: 0, y: 0, radius: 5}
    fmt.Printf("Circle area %f", circle.area())
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

// Close over a variable i of upper function to form a closure
func getSequence() func() int {
    i := 0
    return func() int {
        i ++
        return i
    }
}

type Circle struct {
    x, y, radius float64
}

func(circle Circle) area() float64 {
    return math.Pi * circle.radius * circle.radius
}
