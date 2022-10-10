package api

import (
	"context"
	proto "file/api/qvbilam/file/v1"
	"file/enum"
	"file/global"
	"file/utils/alibab/vod"
	"file/utils/alibab/vod/upload"
	"file/validate"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateUploadVideo(ctx *gin.Context) {
	userId := 1
	key := global.ServerConfig.OssConfig.Key
	secret := global.ServerConfig.OssConfig.Secrect

	request := validate.CreateVideoValidate{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		HandleValidateError(ctx, err)
		return
	}

	fileName := request.FileName
	tags := request.Tags
	fileSha1 := request.FileSha1
	fileType := request.FileType
	categoryId := 0
	title := fmt.Sprintf("user:%d-video:%s", userId, fileName)
	desc := fileSha1

	// 验证文件是否存在
	exists, err := global.FileVideoServerClient.Exists(context.Background(), &proto.GetVideoRequest{FileSha1: fileSha1})
	if err != nil {
		HandleGrpcErrorToHttp(ctx, err)
		return
	}
	if exists.IsExists == true {
		ErrorAlreadyExists(ctx, "video file already exists")
		return
	}

	// 生成上传凭证
	client, err := vod.InitVodClient(key, secret)
	if err != nil {
		ErrorInternal(ctx, "init vod client error")
		return
	}
	params := upload.NewCreateUpdateVideoParams(title, desc, fileName, tags, int64(categoryId))
	certificate, err := upload.CreateUploadVideoCertificate(client, params)
	if err != nil {
		ErrorInternal(ctx, "create upload video certificate error")
		return
	}

	// 创建上传中的文件
	rsp, err := global.FileVideoServerClient.Create(context.Background(), &proto.UpdateVideoRequest{
		UserId:      int64(userId),
		BusinessId:  certificate.VideoId,
		Status:      enum.FileStatusWait,
		ContentType: fileType,
		Sha1:        fileSha1,
		Channel:     enum.FileChannelOss,
	})
	if err != nil {
		HandleGrpcErrorToHttp(ctx, err)
		return
	}

	SuccessNotMessage(ctx, gin.H{
		"id":          rsp.Id,
		"certificate": certificate,
	})
}

// Play 视频点播相关
func Play(ctx *gin.Context) {

}
