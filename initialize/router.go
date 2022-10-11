package initialize

import (
	"file/global"
	"file/middleware"
	customizeRouter "file/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"home":    global.ServerConfig.Home,
			"name":    global.ServerConfig.Name,
			"tags":    global.ServerConfig.Tags,
			"version": "2.0.0",
		})
	})

	apiRouter := router.Group("/v1")

	// 初始化基础组建路由
	customizeRouter.InitDemoRouter(router)
	customizeRouter.InitOssRouter(apiRouter)
	customizeRouter.InitVideoRouter(apiRouter)
	customizeRouter.InitCallbackRouter(apiRouter)
	return router
}
