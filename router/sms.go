package router

import (
	"github.com/gin-gonic/gin"
	"public/api"
)

func InitSmsRouter(Router *gin.RouterGroup) {
	OssRouter := Router.Group("/public/sms")
	{
		OssRouter.POST("/send", api.SendSms)
		OssRouter.POST("/check", api.CheckSms)
	}
}
