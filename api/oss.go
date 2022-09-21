package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"oss/global"
	"oss/utils/alibab/oss"
	"strings"
)

// Token 获取token
func Token(ctx *gin.Context) {
	o := oss.Oss{
		ExpireTime:  global.ServerConfig.OssConfig.ExpireTime,
		UploadDir:   global.ServerConfig.OssConfig.UploadDir,
		Host:        global.ServerConfig.OssConfig.Host,
		Key:         global.ServerConfig.OssConfig.Key,
		Secrect:     global.ServerConfig.OssConfig.Secrect,
		CallBackUrl: global.ServerConfig.OssConfig.CallbackUrl,
	}
	fmt.Println(123)
	fmt.Printf("%+v", o)

	response := o.GetPolicyToken()
	ctx.Header("Access-Control-Allow-Methods", "POST")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.String(200, response)
}

func Callback(ctx *gin.Context) {
	o := oss.Oss{}
	fmt.Println("\nHandle Post Request ... ")
	// Get PublicKey bytes
	bytePublicKey, err := o.GetPublicKey(ctx)
	if err != nil {
		o.ResponseFailed(ctx)
		return
	}

	// Get Authorization bytes : decode from Base64String
	byteAuthorization, err := o.GetAuthorization(ctx)
	if err != nil {
		o.ResponseFailed(ctx)
		return
	}

	// Get MD5 bytes from Newly Constructed Authorization String.
	byteMD5, bodyStr, err := o.GetMD5FromNewAuthString(ctx)
	if err != nil {
		o.ResponseFailed(ctx)
		return
	}

	decodeUrl, err := url.QueryUnescape(bodyStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(decodeUrl)
	params := make(map[string]string)
	datas := strings.Split(decodeUrl, "&")
	for _, v := range datas {
		sdatas := strings.Split(v, "=")
		fmt.Println(v)
		params[sdatas[0]] = sdatas[1]
	}
	fileName := params["filename"]
	fileUrl := fmt.Sprintf("%s/%s", global.ServerConfig.OssConfig.Host, fileName)

	// verifySignature and response to client
	if o.VerifySignature(bytePublicKey, byteMD5, byteAuthorization) {
		// do something you want accoding to callback_body ...
		ctx.JSON(http.StatusOK, gin.H{
			"url": fileUrl,
		})
		//0.ResponseSuccess(ctx)  // response OK : 200
	} else {
		o.ResponseFailed(ctx) // response FAILED : 400
	}
}
