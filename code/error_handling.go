package main

import (
    "fmt"
    "math"
    "errors"
)

func sqrt(value float64) (float64, error) {
    if value < 0 {
        return 0, errors.New("Math: negative number passed to Sqrt")
    }

    return math.Sqrt(value), nil
}

func main() {
    result, err := sqrt(-1)

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(result)
    }

    result, err = sqrt(4)

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(result)
    }
}
