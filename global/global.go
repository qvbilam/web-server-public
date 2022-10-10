package global

import (
	proto "file/api/qvbilam/file/v1"
	userProto "file/api/qvbilam/user/v1"
	"file/config"
	ut "github.com/go-playground/universal-translator"
)

var (
	ServerConfig config.ServerConfig
	Trans        ut.Translator // 表单验证

	FileVideoServerClient proto.VideoClient
	FileImageServerClient proto.ImageClient
	UserServerClient      userProto.UserClient
)
