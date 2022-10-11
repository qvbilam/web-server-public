package router

import (
	"file/api"
	"github.com/gin-gonic/gin"
)

func InitCallbackRouter(Router *gin.RouterGroup) {
	CallbackRouter := Router.Group("/callback")
	{
		CallbackRouter.POST("/vod", api.AliVideoCallback)
	}
}
