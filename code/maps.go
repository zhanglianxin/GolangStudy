package main

import (
    "fmt"
)

func main() {
    // declare a variable, by default map will be nil
	var countryCapitalMap map[string]string
    // define the map as nil map can not be assigned any value
	countryCapitalMap = make(map[string]string)

    countryCapitalMap["France"] = "Paris"
    countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"

    for country, capital := range countryCapitalMap {
        fmt.Println(country, capital, countryCapitalMap[country])
    }

    // test if entry is present in the map or not
    capital, ok := countryCapitalMap["China"]
    fmt.Printf("%#v %#v\n", capital, ok)

    if ok {
        fmt.Println("Capital of China is", capital)
    } else {
        fmt.Println("Capital of China is not present", capital)
    }

    capital1, ok1 := countryCapitalMap["Japan"]
    fmt.Printf("%#v %#v\n", capital1, ok1)

    if ok1 {
        fmt.Println("Capital of Japan is", capital1)
    } else {
        fmt.Println("Capital of Japan is not present", capital1)
    }

    delete(countryCapitalMap, "Japan")

    capital2, ok2 := countryCapitalMap["Japan"]
    fmt.Printf("%#v %#v\n", capital2, ok2)

    if ok2 {
        fmt.Println("Capital of Japan is", capital2)
    } else {
        fmt.Println("Capital of Japan is not present", capital2)
    }
}
