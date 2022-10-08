package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"oss/global"
	"oss/utils/alibab/vod"
	"oss/utils/alibab/vod/upload"
)

func CreateUploadVideo(ctx *gin.Context) {
	key := global.ServerConfig.OssConfig.Key
	secret := global.ServerConfig.OssConfig.Secrect
	fmt.Printf("key: %s\nsecret: %s", key, secret)

	client, err := vod.InitVodClient(key, secret)
	if err != nil { // todo 处理错误
		panic(any(err))
	}
	params := upload.NewCreateUpdateVideoParams("视频标题", "说明", "文件名称", "", "", 0)
	certificate, err := upload.CreateUploadVideoCertificate(client, params)
	if err != nil { // todo 处理错误
		panic(any(err))
	}
	fmt.Println(certificate)
}

// Play 视频点播相关
func Play(ctx *gin.Context) {

}
