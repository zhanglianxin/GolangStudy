package main

import (
    "fmt"
)

func main() {
    fmt.Println("The range keyword is used in for loop to \n" +
        "iterate over items of an array, slice, channel or map.")

    numbers := []int {0, 1, 2, 4}

    for i, j := range numbers {
        fmt.Println(i, j)
    }

    var str = "abcd"

    for i, j := range str {
        fmt.Println(i, j)
    }

    countryCapitalMap := map[string]string {"France": "Paris", "Italy": "Rome", "Japan": "Tokyo" }

    for country, capital := range countryCapitalMap {
        fmt.Println(country, capital)
    }
}
