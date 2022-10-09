package main

import (
	"file/global"
	"file/initialize"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	// 初始化表单验证
	if err := initialize.InitValidateTran("zh"); err != nil {
		zap.S().Panic("翻译器初始化失败: ", err.Error())
	}
	Router := initialize.InitRouters()

	Name := global.ServerConfig.Name
	Port := global.ServerConfig.Port

	// 启动服务
	go func() {
		zap.S().Infof("%s server start listen: %s:%d", Name, "0.0.0.0", Port)
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
