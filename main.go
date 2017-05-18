package main

import (
	"log"
	"net/http"
	"go2rap/proxy"
	"fmt"
)

func main() {

	startServer()
}

func startServer() {

	fmt.Println("start server...")
	////被代理的服务器
	h := &proxy.HandleProxy{}
	//
	err := http.ListenAndServe(":80", h)

	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
