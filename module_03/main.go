package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 1. 接收客户端 request，并将 request 中带的 header 写入 response header
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			for _, vv := range v {
				fmt.Printf("Header : %s=%s \n", k, vv)
				w.Header().Set(k, vv)
			}
		}
	}

	// 2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	os.Setenv("VERSION", "v1.0")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("Version: %s \n", version)

	// 3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	fmt.Println(r.RemoteAddr)
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("err:", err)
	}
	if net.ParseIP(ip) != nil {
		fmt.Printf("Client IP ===>>%s\n", ip)
		log.Println(ip)
	}
	fmt.Printf("Http Status Code ===>>%d\n", http.StatusOK)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}

func healthz(w http.ResponseWriter, r *http.Request) {
	// 4. 当访问 localhost/healthz 时，应返回 200
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
