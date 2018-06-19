package main

import "fmt"

// 全局变量
var a int = 10

func main() {
    // 本地变量
    var a, b, c int = 10, 20, 0

    fmt.Printf("Value of a in main() = %d\n", a)
    c = sum(a, b)
    fmt.Printf("Value of c in main() = %d\n", c)
}

// 形参
func sum(a, b int) int {
    fmt.Printf("Value of a in sum() = %d\n", a)
    fmt.Printf("Value of b in sum() = %d\n", b)

    return a + b
}

// 默认值
// 本地变量和全局变量初始化时的默认值是 0，指针初始化值是 nil
