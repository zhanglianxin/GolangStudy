package main

import "fmt"

func main() {
    var i, j int

    for i = 1; i < 10; i++ {
        for j = 1; j <=i; j++ {
            fmt.Printf("%d * %d = %d\t", j, i, i * j)
        }
        fmt.Println()
    }
    fmt.Println();

    fmt.Print("Prime: ");
    for i = 2; i < 100; i++ {
        for j = 2; j <= (i / j); j++ {
            if (i % j == 0) {
                break
            }
        }
        if (j > (i / j)) {
            fmt.Printf("%d ", i)
        }
    }
}
