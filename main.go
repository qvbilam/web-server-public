package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"oss/initialize"
	"syscall"
)

func main() {
	initialize.InitLogger()
	//initialize.InitConfig()
	Router := initialize.InitRouters()

	//Name := global.ServerConfig.Name
	//Host := global.ServerConfig.Host
	//Port := global.ServerConfig.Port
	Name := "test"
	Host := "0.0.0.0"
	Port := 9501

	// 启动服务
	go func() {
		zap.S().Infof("%s server start listen: %s:%d", Name, Host, Port)
		if err := Router.Run(fmt.Sprintf(":%d", Port)); err != nil {
			zap.S().Panicf("%s server start error: %s", Name, err.Error())
		}
	}()

	// 接受终止信号(优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// 服务注销
	zap.S().Infof("%s server stop success", Name)
}
