package main

import (
	"ai_invest/conf/config"
	"ai_invest/conf/logs"
	"ai_invest/container"
	"ai_invest/kitex_gen/rpc/aiinvestrpcservice"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	//初始化Container
	c, _ := container.Init()
	logs.InitLogger()   //日志模块初始化
	config.InitConfig() //配置文件初始化

	addr, _ := net.ResolveTCPAddr("tcp", ":8888")
	svr := aiinvestrpcservice.NewServer(
		NewServiceHandler(c),
		server.WithServiceAddr(addr),
	)

	if err := svr.Run(); err != nil {
		log.Fatal(err)
	}
}
