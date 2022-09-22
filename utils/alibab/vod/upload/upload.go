package upload

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
	"strconv"
)

type CreateUpdateVideoParams struct {
	Title       string
	Description string
	FileName    string
	CategoryId  int64
	Cover       string
	Tags        string
}

type RefreshVideoUploadParams struct {
	AliVideoId string
}

type CreateUpdateImageParams struct {
}

// CreateVideoUploadCertificate 创建视频上传凭证
func CreateVideoUploadCertificate(client *vod.Client, params *CreateUpdateVideoParams) (response *vod.CreateUploadVideoResponse, err error) {
	request := vod.CreateCreateUploadVideoRequest()
	request.Title = params.Title
	request.Description = params.Description
	request.FileName = params.FileName
	request.CateId = requests.Integer(strconv.FormatInt(params.CategoryId, 10))
	request.CoverURL = params.Cover
	request.Tags = params.Tags

	request.AcceptFormat = "JSON"
	return client.CreateUploadVideo(request)
}

// RefreshVideoUploadCertificate 刷新视频上传凭证
func RefreshVideoUploadCertificate(client *vod.Client, params *RefreshVideoUploadParams) (response *vod.RefreshUploadVideoResponse, err error) {
	request := vod.CreateRefreshUploadVideoRequest()
	request.VideoId = params.AliVideoId
	request.AcceptFormat = "JSON"

	return client.RefreshUploadVideo(request)
}

// CreateUploadImageCertificate 创建图片上传凭证
func CreateUploadImageCertificate(client *vod.Client) (response *vod.CreateUploadImageResponse, err error) {
	request := vod.CreateCreateUploadImageRequest()

	request.ImageType = "cover"
	request.ImageExt = "jpg"
	request.Title = "Sample Image Title"
	request.CateId = "-1"
	request.Tags = "tag1,tag2"

	request.AcceptFormat = "JSON"
	return client.CreateUploadImage(request)
}
