package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

type RequestParams struct {
	Headers map[string]interface{}
	Body    map[string]interface{}
}

func AliVideoCallback(ctx *gin.Context) {
	// 获取参数
	//r := RequestParams{}
	fmt.Println("========== header ==========")
	for k, v := range ctx.Request.Header {
		//r.Headers[k] = v
		fmt.Printf("header-%s: %s\n", k, v)
	}

	fmt.Println("========== form ==========")
	for k, v := range ctx.Request.PostForm {
		fmt.Printf("param-%s: %s\n", k, v)
	}

	fmt.Println("========== query ==========")
	for k, v := range ctx.Request.URL.Query() {
		fmt.Printf("query-%s: %s\n", k, v)
	}

	req, _ := ctx.GetRawData()
	fmt.Println(json.Marshal(req))
	return

	// 写入文件
	file, err := os.OpenFile("./log/test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0)
	if err != nil {
		fmt.Printf("write log err: %s\n", err)
		return
	}
	log.SetOutput(file)

	params := ctx.Params
	//headers := ctx.Header

	log.Printf("%s param: %s, header: %s\n", time.Now(), params)
}
