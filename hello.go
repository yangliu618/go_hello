package main

import "fmt"
import "net/http"

func main() {
    fmt.Println("Hello, World!")
    http.HandleFunc("/",hello)
    http.ListenAndServe(":8080",nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello!"))
}

func init() {
    fmt.Println("init first ")
}
