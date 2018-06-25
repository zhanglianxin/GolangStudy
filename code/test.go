package main

import "fmt"

// 多返回值函数
func getName() (firstName, middleName, lastName, nickName string) {
    firstName = "May"
    middleName = "M"
    lastName = "Chen"
    nickName = "Babe"
    return
}

func main() {
    a, b, c, d := getName()
    fmt.Println(a, b, c, d)

    str := "Hello,世界"
    for i, ch := range str {
        fmt.Println(i, ch) // ch的类型为rune
    }
}
