package category

import "github.com/aliyun/alibaba-cloud-sdk-go/services/vod"

// CreateCategory 创建分类
func CreateCategory(client *vod.Client) (response *vod.AddCategoryResponse, err error) {
	request := vod.CreateAddCategoryRequest()
	request.CateName = "sample category name"
	request.ParentId = "-1"
	request.AcceptFormat = "JSON"

	return client.AddCategory(request)
}

// UpdateCategory 修改分类
func UpdateCategory(client *vod.Client) (response *vod.UpdateCategoryResponse, err error) {
	request := vod.CreateUpdateCategoryRequest()
	request.CateId = "<CateId>"
	request.CateName = "sample category name"
	request.AcceptFormat = "JSON"

	return client.UpdateCategory(request)
}

// DeleteCategory 删除分类
func DeleteCategory(client *vod.Client) (response *vod.DeleteCategoryResponse, err error) {
	request := vod.CreateDeleteCategoryRequest()
	request.CateId = "<CateId>"
	request.AcceptFormat = "JSON"

	return client.DeleteCategory(request)
}

// GetCategories 获取分类及其自分类
func GetCategories(client *vod.Client) (response *vod.GetCategoriesResponse, err error) {
	request := vod.CreateGetCategoriesRequest()
	request.CateId = "<CateId>"
	request.PageNo = "1"
	request.PageSize = "20"
	request.AcceptFormat = "JSON"

	return client.GetCategories(request)
}
