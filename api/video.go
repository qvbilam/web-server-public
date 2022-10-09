package api

import (
	"file/global"
	"file/utils/alibab/vod"
	"file/utils/alibab/vod/upload"
	"file/validate"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateUploadVideo(ctx *gin.Context) {
	key := global.ServerConfig.OssConfig.Key
	secret := global.ServerConfig.OssConfig.Secrect
	fmt.Printf("key: %s\nsecret: %s", key, secret)

	request := validate.CreateVideoValidate{}
	if err := ctx.BindQuery(&request); err != nil {
		HandleValidateError(ctx, err)
		return
	}

	//title := request.Title
	//desc := request.Desc
	fileName := request.FileName
	cover := request.Cover
	tags := request.Tags
	categoryId := 0

	client, err := vod.InitVodClient(key, secret)
	if err != nil {
		ErrorInternal(ctx, "init vod client error")
		return
	}
	params := upload.NewCreateUpdateVideoParams(fileName, fileName, fileName, cover, tags, int64(categoryId))
	certificate, err := upload.CreateUploadVideoCertificate(client, params)
	if err != nil {
		ErrorInternal(ctx, "create upload video certificate error")
		return
	}

	SuccessNotMessage(ctx, gin.H{
		"id":          1,
		"certificate": certificate,
	})
}

// Play 视频点播相关
func Play(ctx *gin.Context) {

}
