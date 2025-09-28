package main

import (
    "fmt"
    "net/http"
)

// 处理 /hello 请求
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "你好，发！这是你的第一个后端接口。")
}

func main() {
    // 注册路由
    http.HandleFunc("/hello", helloHandler)

    // 启动服务器
    fmt.Println("服务器启动：http://localhost:8080/hello")
    http.ListenAndServe(":8080", nil)
}
