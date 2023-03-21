package router

import (
	"github.com/gin-gonic/gin"
	"public/api"
)

func InitCallbackRouter(Router *gin.RouterGroup) {
	CallbackRouter := Router.Group("/public/callback")
	{
		CallbackRouter.POST("/vod", api.AliVideoCallback)
	}
}
