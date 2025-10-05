package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你好，王发！更新你的第一个后端接口。")
}

func main() {
	// Render 会自动提供 PORT 环境变量
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // 本地调试用
	}

	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server listening on port " + port)
	http.ListenAndServe(":"+port, nil)
}

