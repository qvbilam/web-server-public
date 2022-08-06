package router

import (
	"github.com/gin-gonic/gin"
	"oss/api"
)

func InitOssRouter(Router *gin.RouterGroup) {
	OssRouter := Router.Group("")
	{
		OssRouter.GET("/token", api.Token)
		OssRouter.POST("/callback", api.Callback)
	}
}
