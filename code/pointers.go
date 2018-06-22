package main

import (
    "fmt"
)

func main() {
    var a int = 20
    var ip, ptr *int

    ip = &a // 将普通变量的地址赋值给指针变量

    // 访问普通变量的地址
    fmt.Printf("Address of a variable: %x\n", &a)
    // 访问指针变量存储的（普通变量的地址）值
    fmt.Printf("Address stored in ip variable: %x\n", ip)
    // 访问指针变量所指向的变量的值
    fmt.Printf("Value of *ip variable: %d\n", *ip)
    // The nil pointer is a constant with a value of zero defined in several standard libraries.
    fmt.Printf("Address 0 == nil ? %t\n", ptr == nil)

    fmt.Println()

    const MAX int = 3
    arr := []int{10, 100, 200}
    var ptr1 [MAX]*int // 指针数组

    for i := 0; i < MAX; i++ {
        ptr1[i] = &arr[i]
        fmt.Printf("Value of a [%d] = %d\n", i, *ptr1[i])
    }

    fmt.Println()

    var pptr **int // 指向指针的指针

    ptr = &a
    pptr = &ptr

    fmt.Printf("Value of a = %d\n", a)
    fmt.Printf("Value available at *ptr = %d\n", *ptr)
    fmt.Printf("Value available at **pptr = %d\n", **pptr)

    fmt.Println()

    var x, y int = 100, 200
    fmt.Printf("Before swap, x = %d, y = %d\n", x, y)
    swap(&x, &y) // 函数调用时传递指针
    fmt.Printf("After swap, x = %d, y = %d\n", x, y)
}

// 声明函数的参数为指针类型
func swap(x *int, y *int) {
    var temp int
    temp = *x
    *x = *y
    *y = temp
}
