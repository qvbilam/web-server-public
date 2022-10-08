package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitDemoRouter(Router *gin.Engine) {

	// demo
	Router.LoadHTMLFiles("static/demo/index.html", "static/demo/sts-upload.html", "static/demo/upload-auth.html", "static/demo/video-player.html")
	// 配置静态文件夹路径 第一个参数是api，第二个是文件夹路径
	Router.StaticFS("/static/demo", http.Dir(fmt.Sprintf("static/demo")))

	// oss
	Router.GET("/demo/upload/oss", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "posts/index",
		})
	})

	Router.GET("/demo/upload/video", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.HTML(http.StatusOK, "upload-auth.html", gin.H{
			"title": "posts/upload-auth.html",
		})
	})
}
