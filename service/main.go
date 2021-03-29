package main

import (
	"fmt"
	"goTaxi/bootstrap"
	"goTaxi/service/controller"
	"net"
)

func init() {
	controller.InitCache()
	controller.LoopOrder = make(chan *controller.UserConnOrder, 4396)
}

func main() {
	bootstrap.Start()
	//循环抢单
	go  controller.LoopOrders()
	// 建立 tcp 服务
	listen, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}
	for {
		// 等待客户端建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		// 启动一个单独的 goroutine 去处理连接
		go controller.InformationAnalysis(&conn)
	}
}
