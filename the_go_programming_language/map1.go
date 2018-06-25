package main

import (
    "fmt"
)

type PersonInfo struct {
    ID, Name, Address string
}

func main() {
    var personDB map[string] PersonInfo
    personDB = make(map[string] PersonInfo)

    personDB["zhanglianxin"] = PersonInfo{"1", "zhanglianxin", "CN"}
    personDB["xiaoming"] = PersonInfo{"2", "xiaoming", "CN"}

    person, ok := personDB["zhanglianxin"]

    if ok {
        fmt.Printf("Found person %#v", person)
    } else {
        fmt.Println("Did not find it")
    }
}
