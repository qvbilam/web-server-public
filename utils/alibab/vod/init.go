package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
)

func initConfig() *sdk.Config {
	config := sdk.NewConfig()
	config.AutoRetry = true     // 失败是否自动重试
	config.MaxRetryTime = 3     // 最大重试次数
	config.Timeout = 3000000000 // 连接超时，单位：纳秒；默认为3秒
	return config
}

// InitVodClient 初始化
func InitVodClient(accessKeyId string, accessKeySecret string) (client *vod.Client, err error) {

	// 点播服务接入地域
	regionId := "cn-shanghai"

	// 创建授权对象
	credential := &credentials.AccessKeyCredential{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}

	// 自定义config
	config := initConfig()

	// 创建vodClient实例
	return vod.NewClientWithOptions(regionId, config, credential)
}

// InitSTSVodClient 初始化
func InitSTSVodClient(accessKeyId string, accessKeySecret string, securityToken string) (client *vod.Client, err error) {
	// 点播服务接入地域
	regionId := "cn-shanghai"

	// 创建授权对象
	credential := &credentials.StsTokenCredential{
		AccessKeyId:       accessKeyId,
		AccessKeySecret:   accessKeySecret,
		AccessKeyStsToken: securityToken,
	}

	// 连接设置
	config := initConfig()

	// 创建vodClient实例
	return vod.NewClientWithOptions(regionId, config, credential)
}
