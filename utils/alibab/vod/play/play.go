package play

import "github.com/aliyun/alibaba-cloud-sdk-go/services/vod"

// GetVideoPlayCertificate 创建播放凭证
func GetVideoPlayCertificate(client *vod.Client, videoId string) (response *vod.GetVideoPlayAuthResponse, err error) {
	request := vod.CreateGetVideoPlayAuthRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"

	return client.GetVideoPlayAuth(request)
}

// GetVideoPlayUrl 获取视频播放地址
func GetVideoPlayUrl(client *vod.Client, videoId string) (response *vod.GetPlayInfoResponse, err error) {
	request := vod.CreateGetPlayInfoRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"

	return client.GetPlayInfo(request)
}
