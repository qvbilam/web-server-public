package router

import (
	"file/api"
	"github.com/gin-gonic/gin"
)

func InitVideoRouter(Router *gin.RouterGroup) {
	VideoRouter := Router.Group("video")
	{
		VideoRouter.GET("/:id", api.GetVideoDetail)             // 获取视频详情
		VideoRouter.POST("/certificate", api.CreateUploadVideo) // 创建上传凭证
		VideoRouter.PUT("/certificate", api.RefreshUploadVideo) // 更新上传凭证
	}
}
