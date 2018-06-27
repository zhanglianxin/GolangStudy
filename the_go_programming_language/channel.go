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

// 虽然 select 机制不是专为超时而设计的，却能很方便地解决超时问题。
// 因为 select 的特点是只要其中一个 case 已经完成，
// 程序就会继续往下执行，而不会考虑其他 case 的情况。

// 使用 select 机制可以避免永久等待的问题，因为程序会在 timeout 中获取到一个
// 数据后继续执行，无论对 ch 的读取是否还处于等待状态，从而达成 1 秒超时的效果
func timeoutHandle(ch chan int) {
    timeout := make(chan bool, 1)
    go func() {
        time.Sleep(1e9)
        timeout <- true
    }()

    select {
        case <- ch:
            // 从 ch 中读到数据
        case <- timeout:
            // 一直没有从 ch 中读取到数据，但从 timeout 中读取到了数据
    }
}
