package router

import (
	"github.com/gin-gonic/gin"
	"public/api"
)

func InitVideoRouter(Router *gin.RouterGroup) {
	VideoRouter := Router.Group("video")
	{
		VideoRouter.GET("/:id", api.GetVideoDetail)                      // 获取视频详情
		VideoRouter.GET("/:id/play/certificate", api.GetPlayCertificate) // 获取播放凭证
		VideoRouter.POST("/upload/certificate", api.CreateUploadVideo)   // 创建上传凭证
		VideoRouter.PUT("/upload/certificate", api.RefreshUploadVideo)   // 更新上传凭证
	}
}
