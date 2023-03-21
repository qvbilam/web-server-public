package api

import (
	"context"
	"github.com/gin-gonic/gin"
	proto "public/api/qvbilam/public/v1"
	"public/enum"
	"public/global"
	"public/utils"
	"public/validate"
)

func SendSms(ctx *gin.Context) {
	request := validate.SendSmsRequest{}
	if err := ctx.Bind(&request); err != nil {
		HandleValidateError(ctx, err)
		return
	}

	allowTypes := []interface{}{enum.SmsTypeLogin, enum.SmsTypeLogout, enum.SmsTypePassword}
	if !utils.InArray(request.Type, allowTypes) {
		Error(ctx, "参数错误")
		return
	}

	if _, err := global.SmsServerClient.Send(context.Background(), &proto.SendSmsRequest{
		Mobile:   request.Mobile,
		Type:     request.Type,
		ClientIP: GetClientIP(ctx),
	}); err != nil {
		HandleGrpcErrorToHttp(ctx, err)
		return
	}

	SuccessNotContent(ctx)
}

func CheckSms(ctx *gin.Context) {
	request := validate.CheckSmsRequest{}
	if err := ctx.Bind(&request); err != nil {
		HandleValidateError(ctx, err)
		return
	}

	allowTypes := []interface{}{enum.SmsTypeLogin, enum.SmsTypeLogout, enum.SmsTypePassword}
	if !utils.InArray(request.Type, allowTypes) {
		Error(ctx, "参数错误")
		return
	}

	if _, err := global.SmsServerClient.Check(context.Background(), &proto.CheckSmsRequest{
		Mobile: request.Mobile,
		Type:   request.Type,
		Code:   request.Code,
	}); err != nil {
		HandleGrpcErrorToHttp(ctx, err)
		return
	}

	SuccessNotContent(ctx)
}
