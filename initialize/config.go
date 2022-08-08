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
	portString := *flag.String("port", "9501", "server port")
	port, _ := strconv.Atoi(portString)
	ossExpireTimeString := *flag.String("oss_expire_time", "300", "oss expire time")
	ossExpireTime, _ := strconv.Atoi(ossExpireTimeString)

	global.ServerConfig.Name = *flag.String("name", "default-oss-web-server", "server name")
	//global.ServerConfig.Host = *flag.String("host", "0.0.0.0", "server host")
	global.ServerConfig.Port = int64(port)
	global.ServerConfig.OssConfig.Key = *flag.String("oss_key", "", "oss key")
	global.ServerConfig.OssConfig.Secrect = *flag.String("oss_secrect", "", "oss secret")
	global.ServerConfig.OssConfig.Host = *flag.String("oss_host", "", "oss host")
	global.ServerConfig.OssConfig.CallbackUrl = *flag.String("oss_callback_url", "", "oss callback")
	global.ServerConfig.OssConfig.UploadDir = *flag.String("oss_upload_dir", "", "oss upload dir")
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
