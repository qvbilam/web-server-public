package initialize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/namsral/flag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"oss/global"
)

func InitConfig() {
	//initViperConfig()
	InitEnvConfig()
}

func InitEnvConfig() {
	global.ServerConfig.Name = *flag.String("name", "default-oss-web-server", "server name")
	global.ServerConfig.Host = *flag.String("name", "0.0.0.0", "server host")
	global.ServerConfig.Port = *flag.Int64("name", 9501, "server port")
	global.ServerConfig.OssConfig.Key = *flag.String("name", "", "oss key")
	global.ServerConfig.OssConfig.Secrect = *flag.String("name", "", "oss secret")
	global.ServerConfig.OssConfig.Host = *flag.String("name", "", "oss host")
	global.ServerConfig.OssConfig.CallbackUrl = *flag.String("name", "", "oss callback")
	global.ServerConfig.OssConfig.UploadDir = *flag.String("name", "", "oss upload dir")
	global.ServerConfig.OssConfig.ExpireTime = *flag.Int64("name", 300, "oss expire time")
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
