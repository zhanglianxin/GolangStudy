package main

import (
    "fmt"
)

func main() {
    // declare it as an array without specifying its size
    var numbers []int
    // use `make` function to create a slice
    var numbers1 = make([]int, 3, 5)

    printSlice(numbers)
    printSlice(numbers1)

    // nil slice
    if numbers == nil {
        fmt.Printf("%#v\n", numbers)
    }

    fmt.Println()

    numbers2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
    printSlice(numbers2)

    // subslicing
    fmt.Println("[1:4] == ", numbers2[1:4])
    fmt.Println("[:3]  == ", numbers2[:3])
    fmt.Println("[4:]  == ", numbers2[4:])

    fmt.Println()

    var numbers3 = []int{}
    printSlice(numbers3)

    // append() increase the capacity of a slice
    numbers3 = append(numbers3, 0)
    printSlice(numbers3)
    numbers3 = append(numbers3, 1, 2)
    printSlice(numbers3)
    numbers3 = append(numbers3, 3, 4)
    printSlice(numbers3)

    // copy() the contents of a source slice are copied to a destination slice
    numbers4 := make([]int, len(numbers3), cap(numbers3) * 2)
    copy(numbers3, numbers4)
    printSlice(numbers4)
}

func printSlice(x []int) {
    fmt.Printf("len = %d, cap = %d, slice = %v\n", len(x), cap(x), x)
}
