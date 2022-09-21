package upload

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
)

// CreateVideoUploadCertificate 创建视频上传凭证
func CreateVideoUploadCertificate(client *vod.Client) (response *vod.CreateUploadVideoResponse, err error) {
	request := vod.CreateCreateUploadVideoRequest()
	request.Title = "Sample Video Title"
	request.Description = "Sample Description"
	request.FileName = "/opt/video/sample/video_file.mp4"
	//request.CateId = "-1"
	request.CoverURL = "http://192.168.0.0/16/tps/TB1qnJ1PVXXXXXCXXXXXXXXXXXX-700-700.png"
	request.Tags = "tag1,tag2"

	request.AcceptFormat = "JSON"
	return client.CreateUploadVideo(request)
}

// RefreshVideoUploadCertificate 刷新视频上传凭证
func RefreshVideoUploadCertificate(client *vod.Client) (response *vod.RefreshUploadVideoResponse, err error) {
	request := vod.CreateRefreshUploadVideoRequest()
	request.VideoId = "6657f89a86fa4f76a295ae95636e****"
	request.AcceptFormat = "JSON"

	return client.RefreshUploadVideo(request)
}

// CreateUploadImageCertificate 创建图片上传凭证
func CreateUploadImageCertificate(client *vod.Client) (response *vod.CreateUploadImageResponse, err error) {
	request := vod.CreateCreateUploadImageRequest()

	request.ImageType = "cover"
	request.ImageExt = "jpg"
	request.Title = "Sample Image Title"
	//request.CateId = "-1"
	request.Tags = "tag1,tag2"

	request.AcceptFormat = "JSON"
	return client.CreateUploadImage(request)
}
