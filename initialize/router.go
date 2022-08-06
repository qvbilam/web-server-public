package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"oss/middleware"
	ossRouter "oss/router"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	// demo
	router.LoadHTMLFiles(fmt.Sprintf("static/demo/index.html"))
	// 配置静态文件夹路径 第一个参数是api，第二个是文件夹路径
	router.StaticFS("/static/demo", http.Dir(fmt.Sprintf("static/demo")))

	router.GET("/demo", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "posts/index",
		})
	})

	apiRouter := router.Group("/v1")

	// 初始化基础组建路由
	ossRouter.InitOssRouter(apiRouter)
	return router
}
