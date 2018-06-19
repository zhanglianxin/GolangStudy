package main

import (
    "fmt"
)

// 数组，存储相同类型的固定大小的顺序集合
func main() {
    var balance [10]float32

    var balance1 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
    var balance2 = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}

    fmt.Println(balance, balance1, balance2)

    var n [10]int
    var i, j int

    for i = 0; i < 10; i++ {
        n[i] = i + 100
    }

    for j = 0; j < 10; j++ {
        fmt.Printf("Element [%d] = %d\n", j, n[j])
    }

    fmt.Println("---------")

    var balance3 = []int{1000, 2, 3, 17, 50}
    // Passing arrays to functions
    var avg float32 = getAverage(balance3)
    fmt.Println(avg)
}

func getAverage(arr []int) float32 {
    var i, size, sum int

    size = len(arr)

    for i = 0; i < size; i++ {
        sum += arr[i]
    }

    return float32(sum / size)
}
