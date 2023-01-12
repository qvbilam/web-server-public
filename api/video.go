package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	proto "public/api/qvbilam/public/v1"
	"public/enum"
	"public/global"
	"public/utils/alibab/vod"
	"public/utils/alibab/vod/play"
	"public/utils/alibab/vod/upload"
	"public/validate"
	"strconv"
)

type videoCertificateResponse struct {
	Id          int64       `json:"id"`
	Certificate interface{} `json:"certificate"`
}

type videoDetailResponse struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"user_id"`
	BusinessId  string `json:"business_id"`
	Url         string `json:"url"`
	Channel     string `json:"channel"`
	ContentType string `json:"content_type"`
	Status      string `json:"status"`
	Duration    int64  `json:"duration"`
	Size        int64  `json:"size"`
	Expand      string `json:"expand"`
}

// CreateUploadVideo 创建上传视频凭证
func CreateUploadVideo(ctx *gin.Context) {
	userId := 1 // todo 验证用户id
	key := global.ServerConfig.OssConfig.Key
	secret := global.ServerConfig.OssConfig.Secrect

	request := validate.CreateCertificateVideoValidate{}
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
	videoFile, err := global.FileVideoServerClient.Exists(context.Background(), &proto.GetVideoRequest{FileSha1: fileSha1})
	if err != nil {
		HandleGrpcErrorToHttp(ctx, err)
		return
	}
	fmt.Printf("tmd: %+v\n", videoFile)
	if videoFile.IsExists == true {
		ErrorAlreadyExists(ctx, "video file already exists", &videoDetailResponse{
			Id:         videoFile.Video.Id,
			BusinessId: videoFile.Video.BusinessId,
			Url:        videoFile.Video.Url,
			Channel:    videoFile.Video.Channel,
			Status:     videoFile.Video.Status,
		})
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
		ErrorInternal(ctx, "create upload video certificate error"+err.Error())
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

	SuccessNotMessage(ctx, videoCertificateResponse{
		Id:          rsp.Id,
		Certificate: certificate,
	})
}

// RefreshUploadVideo 刷新视频凭证
func RefreshUploadVideo(ctx *gin.Context) {
	request := validate.RefreshCertificateVideoValidate{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		HandleValidateError(ctx, err)
		return
	}
	businessId := request.BusinessId

	// 验证是否存在
	video, err := global.FileVideoServerClient.GetDetail(context.Background(), &proto.GetVideoRequest{BusinessId: businessId})
	if err != nil {
		HandleGrpcErrorToHttp(ctx, err)
		return
	}
	if video.Id == 0 {
		ErrorNotfound(ctx, "not found video file")
		return
	}

	// 生成上传凭证
	key := global.ServerConfig.OssConfig.Key
	secret := global.ServerConfig.OssConfig.Secrect
	client, err := vod.InitVodClient(key, secret)
	if err != nil {
		ErrorInternal(ctx, "init vod client error")
		return
	}

	certificate, err := upload.RefreshUploadVideoCertificate(client, &upload.RefreshVideoUploadParams{AliVideoId: businessId})
	if err != nil {
		ErrorInternal(ctx, "refresh upload video certificate error"+err.Error())
		return
	}

	SuccessNotMessage(ctx, videoCertificateResponse{
		Id:          video.Id,
		Certificate: certificate,
	})
}

// GetVideoDetail 获取视频信息
func GetVideoDetail(ctx *gin.Context) {
	request := validate.GetFileVideoValidate{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		HandleValidateError(ctx, err)
		return
	}
	paramId := ctx.Param("id")
	id, _ := strconv.Atoi(paramId)

	businessId := request.BusinessId
	fileSha1 := request.FileSha1
	if id == 0 && businessId == "" && fileSha1 == "" {
		Error(ctx, "bad request")
		return
	}

	v, err := global.FileVideoServerClient.GetDetail(context.Background(), &proto.GetVideoRequest{
		Id:         int64(id),
		BusinessId: businessId,
		FileSha1:   fileSha1,
	})
	if err != nil {
		HandleGrpcErrorToHttp(ctx, err)
		return
	}

	SuccessNotMessage(ctx, videoDetailResponse{
		Id:          v.Id,
		UserId:      v.UserId,
		BusinessId:  v.BusinessId,
		Url:         v.Url,
		Duration:    v.Duration,
		Size:        v.Size,
		Channel:     v.Channel,
		ContentType: v.ContentType,
		Status:      v.Status,
		Expand:      v.Expand,
	})
}

// GetPlayCertificate 获取播放凭证
func GetPlayCertificate(ctx *gin.Context) {
	idString := ctx.Query("id")
	id, _ := strconv.Atoi(idString)
	v, err := global.FileVideoServerClient.GetDetail(context.Background(), &proto.GetVideoRequest{Id: int64(id)})
	if err != nil {
		HandleGrpcErrorToHttp(ctx, err)
		return
	}
	businessId := v.BusinessId

	key := global.ServerConfig.OssConfig.Key
	secret := global.ServerConfig.OssConfig.Secrect
	client, err := vod.InitVodClient(key, secret)
	if err != nil {
		ErrorInternal(ctx, "init vod client error")
		return
	}
	certificate, err := play.GetVideoPlayCertificate(client, businessId)
	if err != nil {
		Error(ctx, "get video play error")
		return
	}

	SuccessNotMessage(ctx, gin.H{
		"id":          id,
		"certificate": certificate,
	})
}
