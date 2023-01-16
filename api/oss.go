package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
	proto "public/api/qvbilam/public/v1"
	"public/global"
	"public/utils/alibab/oss"
	"public/validate"
	"strings"
)

// Token 获取token
func Token(ctx *gin.Context) {

	request := validate.CreateOss{}
	if err := ctx.Bind(&request); err != nil {
		HandleValidateError(ctx, err)
		return
	}
	// 判断文件是否存在
	exists, err := global.FileServerClient.Exists(context.Background(), &proto.FileDetailRequest{
		Sha1: request.FileSha1,
	})
	if err != nil {
		HandleGrpcErrorToHttp(ctx, err)
		return
	}

	if exists.IsExists {
		ErrorAlreadyExists(ctx, "文件已存在", gin.H{
			"id":           exists.File.Id,
			"url":          exists.File.Url,
			"file_sha1":    exists.File.Sha1,
			"size":         exists.File.Size,
			"content_type": exists.File.ContentType,
			"channel":      exists.File.Channel,
		})
		return
	}

	o := oss.Oss{
		ExpireTime:  global.ServerConfig.OssConfig.ExpireTime,
		UploadDir:   global.ServerConfig.OssConfig.UploadDir,
		Host:        global.ServerConfig.OssConfig.Host,
		Key:         global.ServerConfig.OssConfig.Key,
		Secrect:     global.ServerConfig.OssConfig.Secrect,
		CallBackUrl: global.ServerConfig.OssConfig.CallbackUrl,
	}

	response, err := o.GetPolicyToken()
	if err != nil {
		Error(ctx, err.Error())
		return
	}

	SuccessNotMessage(ctx, response)
}

func Callback(ctx *gin.Context) {
	//requestLog(ctx) // 记录请求日志

	o := oss.Oss{}
	// Get PublicKey bytes
	bytePublicKey, err := o.GetPublicKey(ctx)
	if err != nil {
		Error(ctx, err.Error())
		return
	}

	// Get Authorization bytes : decode from Base64String
	byteAuthorization, err := o.GetAuthorization(ctx)
	if err != nil {
		Error(ctx, err.Error())
		return
	}

	// Get MD5 bytes from Newly Constructed Authorization String.
	byteMD5, bodyStr, err := o.GetMD5FromNewAuthString(ctx)
	if err != nil {
		Error(ctx, err.Error())
		return
	}

	decodeUrl, err := url.QueryUnescape(bodyStr)
	if err != nil {
		Error(ctx, err.Error())
		return
	}
	//fmt.Println(decodeUrl)
	params := make(map[string]string)
	ds := strings.Split(decodeUrl, "&")
	for _, v := range ds {
		sds := strings.Split(v, "=")
		fmt.Println(v)
		params[sds[0]] = sds[1]
	}
	fileName := params["filename"]
	fileUrl := fmt.Sprintf("%s/%s", global.ServerConfig.OssConfig.Host, fileName)

	// verifySignature and response to client
	if !o.VerifySignature(bytePublicKey, byteMD5, byteAuthorization) {
		// do something you want accoding to callback_body ...

		Error(ctx, "verifySignature fail")
		return
	}

	SuccessNotMessage(ctx, gin.H{
		"url": fileUrl,
	})

}
