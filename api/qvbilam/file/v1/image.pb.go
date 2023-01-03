// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.19.4
// source: image.proto

package fileV1

import (
	v1 "file/api/qvbilam/page/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	BusinessId  string `protobuf:"bytes,3,opt,name=businessId,proto3" json:"businessId,omitempty"`
	Url         string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Size        int64  `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Width       int64  `protobuf:"varint,6,opt,name=width,proto3" json:"width,omitempty"`
	Height      int64  `protobuf:"varint,7,opt,name=height,proto3" json:"height,omitempty"`
	Status      string `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	ContentType string `protobuf:"bytes,9,opt,name=contentType,proto3" json:"contentType,omitempty"`
	Expand      string `protobuf:"bytes,10,opt,name=expand,proto3" json:"expand,omitempty"`
	CreatedAt   string `protobuf:"bytes,11,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Channel     string `protobuf:"bytes,12,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *ImageResponse) Reset() {
	*x = ImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageResponse) ProtoMessage() {}

func (x *ImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageResponse.ProtoReflect.Descriptor instead.
func (*ImageResponse) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{0}
}

func (x *ImageResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ImageResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ImageResponse) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *ImageResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ImageResponse) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ImageResponse) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ImageResponse) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ImageResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ImageResponse) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *ImageResponse) GetExpand() string {
	if x != nil {
		return x.Expand
	}
	return ""
}

func (x *ImageResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ImageResponse) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

type ImagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total  int64             `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Images []*ImagesResponse `protobuf:"bytes,2,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *ImagesResponse) Reset() {
	*x = ImagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImagesResponse) ProtoMessage() {}

func (x *ImagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImagesResponse.ProtoReflect.Descriptor instead.
func (*ImagesResponse) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{1}
}

func (x *ImagesResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ImagesResponse) GetImages() []*ImagesResponse {
	if x != nil {
		return x.Images
	}
	return nil
}

type UpdateImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	BusinessId  string `protobuf:"bytes,3,opt,name=businessId,proto3" json:"businessId,omitempty"`
	Url         string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Size        int64  `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Width       int64  `protobuf:"varint,6,opt,name=width,proto3" json:"width,omitempty"`
	Height      int64  `protobuf:"varint,7,opt,name=height,proto3" json:"height,omitempty"`
	Status      string `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	ContentType string `protobuf:"bytes,9,opt,name=contentType,proto3" json:"contentType,omitempty"`
	Expand      string `protobuf:"bytes,10,opt,name=expand,proto3" json:"expand,omitempty"`
	Sha1        string `protobuf:"bytes,11,opt,name=sha1,proto3" json:"sha1,omitempty"`
	Channel     string `protobuf:"bytes,12,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *UpdateImageRequest) Reset() {
	*x = UpdateImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateImageRequest) ProtoMessage() {}

func (x *UpdateImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateImageRequest.ProtoReflect.Descriptor instead.
func (*UpdateImageRequest) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateImageRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateImageRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateImageRequest) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *UpdateImageRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UpdateImageRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *UpdateImageRequest) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *UpdateImageRequest) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *UpdateImageRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UpdateImageRequest) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *UpdateImageRequest) GetExpand() string {
	if x != nil {
		return x.Expand
	}
	return ""
}

func (x *UpdateImageRequest) GetSha1() string {
	if x != nil {
		return x.Sha1
	}
	return ""
}

func (x *UpdateImageRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

type GetImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BusinessId string `protobuf:"bytes,2,opt,name=businessId,proto3" json:"businessId,omitempty"`
	FileSha1   string `protobuf:"bytes,3,opt,name=fileSha1,proto3" json:"fileSha1,omitempty"`
}

func (x *GetImageRequest) Reset() {
	*x = GetImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageRequest) ProtoMessage() {}

func (x *GetImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageRequest.ProtoReflect.Descriptor instead.
func (*GetImageRequest) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{3}
}

func (x *GetImageRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetImageRequest) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *GetImageRequest) GetFileSha1() string {
	if x != nil {
		return x.FileSha1
	}
	return ""
}

type SearchImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         []int64         `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
	UserId     []int64         `protobuf:"varint,2,rep,packed,name=userId,proto3" json:"userId,omitempty"`
	BusinessId []string        `protobuf:"bytes,3,rep,name=businessId,proto3" json:"businessId,omitempty"`
	FileSha1   string          `protobuf:"bytes,4,opt,name=fileSha1,proto3" json:"fileSha1,omitempty"`
	Page       *v1.PageRequest `protobuf:"bytes,5,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *SearchImageRequest) Reset() {
	*x = SearchImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchImageRequest) ProtoMessage() {}

func (x *SearchImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchImageRequest.ProtoReflect.Descriptor instead.
func (*SearchImageRequest) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{4}
}

func (x *SearchImageRequest) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *SearchImageRequest) GetUserId() []int64 {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *SearchImageRequest) GetBusinessId() []string {
	if x != nil {
		return x.BusinessId
	}
	return nil
}

func (x *SearchImageRequest) GetFileSha1() string {
	if x != nil {
		return x.FileSha1
	}
	return ""
}

func (x *SearchImageRequest) GetPage() *v1.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

type ExistsImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsExists bool           `protobuf:"varint,1,opt,name=isExists,proto3" json:"isExists,omitempty"`
	Image    *ImageResponse `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *ExistsImageResponse) Reset() {
	*x = ExistsImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistsImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistsImageResponse) ProtoMessage() {}

func (x *ExistsImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistsImageResponse.ProtoReflect.Descriptor instead.
func (*ExistsImageResponse) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{5}
}

func (x *ExistsImageResponse) GetIsExists() bool {
	if x != nil {
		return x.IsExists
	}
	return false
}

func (x *ExistsImageResponse) GetImage() *ImageResponse {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_image_proto protoreflect.FileDescriptor

var file_image_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x76, 0x62, 0x69, 0x6c,
	0x61, 0x6d, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x02, 0x0a, 0x0d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x78, 0x70, 0x61, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x59, 0x0a,
	0x0e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x31, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0xb0, 0x02, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x69,
	0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x68, 0x61, 0x31, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x68, 0x61,
	0x31, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x5d, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x31, 0x22, 0xa4, 0x01, 0x0a, 0x12, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x68, 0x61, 0x31, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x68, 0x61, 0x31, 0x12, 0x2a, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x22, 0x61, 0x0a, 0x13, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x32, 0x96, 0x03, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x41,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x50,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3f, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x3f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x1a, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x12, 0x1a, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x21, 0x5a,
	0x1f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x76, 0x62, 0x69, 0x6c, 0x61,
	0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x69, 0x6c, 0x65, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_image_proto_rawDescOnce sync.Once
	file_image_proto_rawDescData = file_image_proto_rawDesc
)

func file_image_proto_rawDescGZIP() []byte {
	file_image_proto_rawDescOnce.Do(func() {
		file_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_image_proto_rawDescData)
	})
	return file_image_proto_rawDescData
}

var file_image_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_image_proto_goTypes = []interface{}{
	(*ImageResponse)(nil),       // 0: filePb.v1.ImageResponse
	(*ImagesResponse)(nil),      // 1: filePb.v1.ImagesResponse
	(*UpdateImageRequest)(nil),  // 2: filePb.v1.UpdateImageRequest
	(*GetImageRequest)(nil),     // 3: filePb.v1.GetImageRequest
	(*SearchImageRequest)(nil),  // 4: filePb.v1.SearchImageRequest
	(*ExistsImageResponse)(nil), // 5: filePb.v1.ExistsImageResponse
	(*v1.PageRequest)(nil),      // 6: pagePb.v1.PageRequest
	(*emptypb.Empty)(nil),       // 7: google.protobuf.Empty
}
var file_image_proto_depIdxs = []int32{
	1, // 0: filePb.v1.ImagesResponse.images:type_name -> filePb.v1.ImagesResponse
	6, // 1: filePb.v1.SearchImageRequest.page:type_name -> pagePb.v1.PageRequest
	0, // 2: filePb.v1.ExistsImageResponse.image:type_name -> filePb.v1.ImageResponse
	2, // 3: filePb.v1.Image.Create:input_type -> filePb.v1.UpdateImageRequest
	2, // 4: filePb.v1.Image.Update:input_type -> filePb.v1.UpdateImageRequest
	2, // 5: filePb.v1.Image.Delete:input_type -> filePb.v1.UpdateImageRequest
	4, // 6: filePb.v1.Image.Get:input_type -> filePb.v1.SearchImageRequest
	3, // 7: filePb.v1.Image.GetDetail:input_type -> filePb.v1.GetImageRequest
	3, // 8: filePb.v1.Image.Exists:input_type -> filePb.v1.GetImageRequest
	0, // 9: filePb.v1.Image.Create:output_type -> filePb.v1.ImageResponse
	7, // 10: filePb.v1.Image.Update:output_type -> google.protobuf.Empty
	7, // 11: filePb.v1.Image.Delete:output_type -> google.protobuf.Empty
	1, // 12: filePb.v1.Image.Get:output_type -> filePb.v1.ImagesResponse
	0, // 13: filePb.v1.Image.GetDetail:output_type -> filePb.v1.ImageResponse
	5, // 14: filePb.v1.Image.Exists:output_type -> filePb.v1.ExistsImageResponse
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_image_proto_init() }
func file_image_proto_init() {
	if File_image_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImagesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_image_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateImageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_image_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_image_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchImageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_image_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistsImageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_image_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_image_proto_goTypes,
		DependencyIndexes: file_image_proto_depIdxs,
		MessageInfos:      file_image_proto_msgTypes,
	}.Build()
	File_image_proto = out.File
	file_image_proto_rawDesc = nil
	file_image_proto_goTypes = nil
	file_image_proto_depIdxs = nil
}
