package initialize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"public/global"
	"strconv"
)

func InitConfig() {
	InitEnvConfig()
	initViperConfig()
}

func InitEnvConfig() {
	serverPort, _ := strconv.Atoi(os.Getenv("PORT"))
	ossExpiredTime, _ := strconv.Atoi(os.Getenv("OSS_EXPIRE_TIME"))
	userServerPort, _ := strconv.Atoi(os.Getenv("USER_SERVER_PORT"))
	publicServerPort, _ := strconv.Atoi(os.Getenv("PUBLIC_SERVER_PORT"))
	// 服务
	global.ServerConfig.Name = os.Getenv("SERVER_NAME")
	global.ServerConfig.Home = "0.0.0.0"
	global.ServerConfig.Port = int64(serverPort)
	// oss
	global.ServerConfig.OssConfig.Key = os.Getenv("OSS_KEY")
	global.ServerConfig.OssConfig.Secret = os.Getenv("OSS_SECRET")
	global.ServerConfig.OssConfig.Host = os.Getenv("OSS_HOST")
	global.ServerConfig.OssConfig.CallbackUrl = os.Getenv("OSS_CALLBACK_URL")
	global.ServerConfig.OssConfig.UploadDir = os.Getenv("OSS_UPLOAD_DIR")
	global.ServerConfig.OssConfig.ExpireTime = int64(ossExpiredTime)
	// user-server
	global.ServerConfig.UserServerConfig.Host = os.Getenv("USER_SERVER_HOST")
	global.ServerConfig.UserServerConfig.Name = os.Getenv("USER_SERVER_NAME")
	global.ServerConfig.UserServerConfig.Port = int64(userServerPort)
	// public-server
	global.ServerConfig.PublicServerConfig.Host = os.Getenv("PUBLIC_SERVER_HOST")
	global.ServerConfig.PublicServerConfig.Name = os.Getenv("PUBLIC_SERVER_NAME")
	global.ServerConfig.PublicServerConfig.Port = int64(publicServerPort)
}

func initViperConfig() {
	file := "config.yaml"
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return
	}

	v := viper.New()
	v.SetConfigFile(file)
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
