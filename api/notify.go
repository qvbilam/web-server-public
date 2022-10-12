package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

type RequestParams struct {
	Headers map[string]interface{}
	Body    map[string]interface{}
}

func AliVideoCallback(ctx *gin.Context) {
	// 获取参数
	r := RequestParams{}
	body := make(map[string]interface{})
	headers := make(map[string]interface{})
	// header
	for k, v := range ctx.Request.Header {
		//r.Headers[k] = v
		//fmt.Printf("header-%s: %s\n", k, v[0])
		headers[k] = v[0]
	}

	// form
	data, _ := ctx.GetRawData()
	var formBody map[string]string
	_ = json.Unmarshal(data, &body)
	for k, v := range formBody {
		//fmt.Printf("form-%s: %s\n", k, v)
		body[k] = v[0]
	}

	// query
	for k, v := range ctx.Request.URL.Query() {
		//fmt.Printf("query-%s: %s\n", k, v)
		body[k] = v[0]
	}

	r.Body = body
	r.Headers = headers
	req, _ := json.Marshal(r)

	// 写入文件
	fileName := fmt.Sprintf("%s-%s.log", "notify-video", time.Now().Format("2006-01-02"))
	filePath := fmt.Sprintf("./log/%s", fileName)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	if err != nil {
		fmt.Printf("write log err: %s\n", err)
		return
	}
	log.SetOutput(file)

	log.SetFlags(log.Ltime | log.Ldate)
	log.Printf(string(req))

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
