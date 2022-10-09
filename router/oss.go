package router

import (
	"file/api"
	"github.com/gin-gonic/gin"
)

func InitOssRouter(Router *gin.RouterGroup) {
	OssRouter := Router.Group("")
	{
		OssRouter.GET("/token", api.Token)
		OssRouter.POST("/callback", api.Callback)
	}
}
