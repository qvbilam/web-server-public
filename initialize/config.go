package initialize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/namsral/flag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"oss/global"
	"strconv"
)

func InitConfig() {
	//initViperConfig()
	InitEnvConfig()
}

func InitEnvConfig() {
	serverPort := flag.String("port", "9501", "server port")
	serverName := flag.String("server_name", "default-oss-web-server", "server name")
	ossKey := flag.String("oss_key", "", "oss key")
	ossSecrect := flag.String("oss_secrect", "", "oss secrect")
	ossHost := flag.String("oss_host", "", "oss host")
	ossCallbackUrl := flag.String("oss_callback_url", "", "oss callback")
	ossUploadDir := flag.String("oss_upload_dir", "", "oss upload dir")
	ossExpireTimeString := flag.String("oss_expire_time", "300", "oss expire time")
	// 必须加, 否则将获取的配置值解析为定义的变量中
	flag.Parse()
	port, _ := strconv.Atoi(*serverPort)
	ossExpireTime, _ := strconv.Atoi(*ossExpireTimeString)

	// 写入全局变量
	global.ServerConfig.Name = *serverName
	global.ServerConfig.Port = int64(port)
	global.ServerConfig.OssConfig.Key = *ossKey
	global.ServerConfig.OssConfig.Secrect = *ossSecrect
	global.ServerConfig.OssConfig.Host = *ossHost
	global.ServerConfig.OssConfig.CallbackUrl = *ossCallbackUrl
	global.ServerConfig.OssConfig.UploadDir = *ossUploadDir
	global.ServerConfig.OssConfig.ExpireTime = int64(ossExpireTime)
}

func initViperConfig() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panicf("获取配置异常: %s", err)
	}
	// 映射配置文件
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		zap.S().Panicf("加载配置异常: %s", err)
	}
	// 动态监听配置
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&global.ServerConfig)
	})
}
