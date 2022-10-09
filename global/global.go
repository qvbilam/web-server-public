package global

import (
	"file/config"
	ut "github.com/go-playground/universal-translator"
)

var (
	ServerConfig config.ServerConfig
	Trans        ut.Translator // 表单验证
)
