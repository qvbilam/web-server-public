package global

import (
	ut "github.com/go-playground/universal-translator"
	proto "public/api/qvbilam/public/v1"
	userProto "public/api/qvbilam/user/v1"
	"public/config"
)

var (
	ServerConfig config.ServerConfig
	Trans        ut.Translator // 表单验证

	FileVideoServerClient proto.VideoClient
	FileServerClient      proto.FileClient
	SmsServerClient       proto.SmsClient
	UserServerClient      userProto.UserClient
)
