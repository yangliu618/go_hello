package main

import "fmt"
import "net/http"

func main() {
    fmt.Println("Hello, World!" + " Jack")
    fmt.Println("1+3=",1+3)
    fmt.Println(true)

    //字符串可以使用 + 来连接
    fmt.Println("go" + "lang")

    //整型和浮点数
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    //布尔类型，你可以使用布尔运算符
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)


    http.HandleFunc("/",hello)
    http.ListenAndServe(":8080",nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello!"))
}

func init() {
    fmt.Println("init first ")
}
