package validate

type CreateOss struct {
	FileSha1 string `form:"file_sha1" json:"file_sha1" binding:"required"`
}