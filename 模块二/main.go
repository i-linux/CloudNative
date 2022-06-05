package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// HTTP 服务器示例
func main() {
	flag.Parse()
	log.Println("Starting http server...")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	// (4) 当访问 localhost/healthz 时，应返回 200
	// io.WriteString(w, "200\n")
	w.Write([]byte("200\n"))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// (1) 接收客户端 request，并将 request 中带的 header 写入 response header
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			w.Header().Add(k, fmt.Sprint(v))
		}
	}

	// (2) 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	w.Header().Add("VERSION", os.Getenv("VERSION"))

	// (3) Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	log.Println("Client IP:", strings.Split(r.RemoteAddr, ":")[0],
		", Return Code:", http.StatusOK)

	w.Write([]byte("Hello World\n"))
}
