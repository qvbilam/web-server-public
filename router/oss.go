package router

import (
	"github.com/gin-gonic/gin"
	"public/api"
)

func InitOssRouter(Router *gin.RouterGroup) {
	OssRouter := Router.Group("/public/oss")
	{
		OssRouter.GET("/token", api.Token)
		OssRouter.POST("/callback", api.Callback)
	}
}
