package main

import (
    "fmt"
)

func main() {
    var i int = 15
    fmt.Printf("factorial of %d is %d\n", i, factorial(i))

    var scope int = 10
    fmt.Print("fibonacci series: ")
    for i := 0; i < scope; i++ {
        fmt.Printf("%d ", fibonaci(i))
    }
}

func factorial(i int) int {
    if i <= 1 {
        return 1
    }
    return i * factorial(i - 1)
}

func fibonaci(i int) (ret int) {
    if i == 0 || i == 1 {
        return i
    }
    return fibonaci(i - 1) + fibonaci(i - 2)
}
