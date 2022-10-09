package router

import (
	"file/api"
	"github.com/gin-gonic/gin"
)

func InitVideoRouter(Router *gin.RouterGroup) {
	VideoRouter := Router.Group("video")
	{
		VideoRouter.POST("/", api.CreateUploadVideo)
	}
}
