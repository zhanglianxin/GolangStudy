package main

import "fmt"

func main() {
    var a int = 4
    var ptr *int

    fmt.Println("Hello, World!")

    /* example of & and * operators */
    ptr = &a /* 'ptr' now contains the address of 'a' */
    fmt.Printf("value of a is %d\n", a)
    fmt.Printf("*ptr is %d.\n", *ptr)
}
