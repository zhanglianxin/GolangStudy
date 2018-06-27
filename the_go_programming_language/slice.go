package main

import (
    "fmt"
)

func main() {
    var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    // 基于数组创建切片
    var mySlice []int = myArray[:5]
    // 直接创建切片
    mySlice1 := make([]int, 5, 10)
    // 直接创建并初始化
    mySlice2 := []int{1, 2, 3, 4, 5}

    fmt.Println("Elements of myArray: ")
    for _, v := range myArray {
        fmt.Printf("%d ", v)
    }

    fmt.Println("\nElements of mySlice: ")
    for _, v := range mySlice {
        fmt.Printf("%d ", v)
    }

    fmt.Println("\nElements of mySlice1: ")
    for _, v := range mySlice1 {
        fmt.Printf("%d ", v)
    }

    fmt.Println("\nElements of mySlice2: ")
    for _, v := range mySlice2 {
        fmt.Printf("%d ",  v)
    }

    fmt.Println()

    sl := make([]int, 1)
    fmt.Printf("%#v\n", sl)
    sl = sl[1:] // 不会超出边界
    fmt.Printf("%#v\n", sl)
}
