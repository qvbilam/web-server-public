package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"oss/global"
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

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"home":    global.ServerConfig.Home,
			"name":    global.ServerConfig.Name,
			"tags":    global.ServerConfig.Tags,
			"version": "2.0.0",
		})
	})

	// demo
	router.LoadHTMLFiles("static/demo/index.html", "static/demo/sts-upload.html", "static/demo/upload-auth.html", "static/demo/video-player.html")
	// 配置静态文件夹路径 第一个参数是api，第二个是文件夹路径
	router.StaticFS("/static/demo", http.Dir(fmt.Sprintf("static/demo")))

	router.GET("/demo", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "posts/index",
		})
	})

	router.GET("/demo/upload", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.HTML(http.StatusOK, "upload-auth.html", gin.H{
			"title": "posts/upload-auth.html",
		})
	})

	apiRouter := router.Group("/v1")

	// 初始化基础组建路由
	ossRouter.InitOssRouter(apiRouter)
	return router
}
