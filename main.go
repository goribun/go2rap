package main

import (
	"log"
	"net/http"
	"go2rap/proxy"
)

func main() {
	startServer()
}

//启动服务
func startServer() {

	log.Println("starting server...")
	//反向代理处理器
	h := &proxy.HandleProxy{}
	//监听80端口
	err := http.ListenAndServe(":80", h)

	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
