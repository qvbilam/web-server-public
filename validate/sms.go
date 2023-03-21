package validate

type SendSmsRequest struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
	Type   string `form:"type" json:"type" binding:"required"`
}

type CheckSmsRequest struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
	Code   string `form:"code" json:"code" binding:"required"`
	Type   string `form:"type" json:"type" binding:"required"`
}
