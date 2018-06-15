package main

import "fmt"

func main() {
    var a int = 10

    fmt.Println("break")
    for a < 20 {
        fmt.Printf("%d ", a)
        a++
        if a > 15 {
            break
        }
    }
    fmt.Println()

    fmt.Println("continue")
    a = 10
    for a < 20 {
        if a == 15 {
            a++
            continue
        }
        fmt.Printf("%d ", a)
        a++
    }
    fmt.Println()

    fmt.Println("goto")
    a = 10
    LOOP: for a < 20 {
        if a == 15 {
            a++
            goto LOOP
        }
        fmt.Printf("%d ", a)
        a++
    }
    fmt.Println()

    fmt.Println("infinite loop")
    a = 10
    for {
        if a == 20 {
            break
        }
        fmt.Printf("%d ", a)
        a++
    }
}
