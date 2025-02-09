package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// 获取请求的域名
	host := r.Host
	fmt.Printf("Received request from domain: %s\n", host)

	// 向客户端响应
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprintf("Domain received: %s", host)))
}

func main() {
	// 设置处理函数
	http.HandleFunc("/", handler)

	// 启动服务器
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
