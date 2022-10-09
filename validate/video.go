package validate

type CreateVideoValidate struct {
	//Title    string `form:"title" json:"title" binding:"required,min=1,max=50"`
	//Desc     string `form:"desc" json:"desc" binding:"required,max=1024"`
	FileName string `form:"file_name" json:"file_name" binding:"required"`
	FileSha1 string `form:"file_sha1" json:"file_sha_1" binding:"required"`
	Cover    string `form:"cover" json:"cover" binding:"required,url"`
	Tags     string `form:"tags" json:"tags" binding:"omitempty"`
}
