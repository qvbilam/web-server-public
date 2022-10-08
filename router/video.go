package router

import (
	"github.com/gin-gonic/gin"
	"oss/api"
)

func InitVideoRouter(Router *gin.RouterGroup) {
	VideoRouter := Router.Group("video")
	{
		VideoRouter.POST("/", api.CreateUploadVideo)
	}
}
