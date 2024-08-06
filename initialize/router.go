package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"public/middleware"
	customizeRouter "public/router"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.GET("/public/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	router.GET("/public/version", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "v1.0.0",
		})
	})

	apiRouter := router.Group("")

	// 初始化基础组建路由
	customizeRouter.InitDemoRouter(router)
	customizeRouter.InitOssRouter(apiRouter)
	customizeRouter.InitVideoRouter(apiRouter)
	customizeRouter.InitCallbackRouter(apiRouter)
	customizeRouter.InitSmsRouter(apiRouter)
	return router
}
