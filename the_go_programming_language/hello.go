package main

import (
    "net/http"
    "io"
    "log"
)

func main() {
    http.HandleFunc("/hello", helloHandler)
    err := http.ListenAndServe(":8090", nil)
    if err != nil {
        log.Fatal("ListenAndService: ", err.Error())
    }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello, world!")
}
