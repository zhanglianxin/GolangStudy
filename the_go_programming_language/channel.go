package main

import (
    "fmt"
    "time"
)

func main() {
    chs := make([]chan int, 10)
    for i := 0; i < 10; i++ {
        chs[i] = make(chan int)
        go Count(chs[i], i)
    }

    for _, ch := range chs {
        fmt.Println(<- ch)
    }

    fmt.Println()

    randomWrite2Channel()
}

func Count(ch chan int, i int) {
    fmt.Println("Counting", i)
    ch <- i
}

// 每个 case 语句里必须是一个 IO 操作
func selectKeyword() {
    /*var chan1, chan2 chan int
    select {
        case <- chan1:
            // 如果 chan1 成功读到数据，则进行该 case 处理语句
        case chan2 <- 1:
            // 如果成功向 chan2 写入数据，则进行该 case 处理语句
        default:
    }*/
}

// 利用 select 随机向 ch 中写入一个 0 或者 1
func randomWrite2Channel() {
    ch := make(chan int, 1)
    j := 0
    for {
        select {
            case ch <- 0:
            case ch <- 1:
        }
        i := <-ch
        fmt.Println("Value received:", i)
        j++
        if j == 20 {
            return
        }
    }
}

func timeoutHandle() {
    timeout := make(chan bool, 1)
    go func() {
        time.Sleep(le9)
        timeout <- true
    }()

    select {
        case <-ch:
            // 从 ch 中读到数据
        case <-time:
            // 一直没有从 ch 中读取到数据，但从 timeout 中读取到了数据
    }
}
