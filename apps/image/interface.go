package image

import (
	"context"
	"strings"
)

// 模块名称
const (
	AppName = "image"
)

// Image功能接口
type Service interface {
	CreateImage(context.Context, *CreateImageRequest) (*CreateImageResponse, error)
	CreateEditImage(context.Context, *CreateEditImageRequest) (*CreateImageResponse, error)
}

// CreateImageRequest结构体
type CreateImageRequest struct {
	Prompt         string `json:"prompt"`
	Model          string `json:"model"`
	N              int    `json:"n"`
	Quality        string `json:"quality"`
	ResponseFormat string `json:"response_format"`
	Size           string `json:"size"`
	Style          string `json:"style"`
	User           string `json:"user"`
}

// CreateImageRequest结构体构造函数
func NewCreateImageRequest() *CreateImageRequest {
	return &CreateImageRequest{}
}

// CreateEditImageRequest结构体
type CreateEditImageRequest struct {
	ImageFilePath  string `json:"image_file_path"`
	ImageFileName  string `json:"image_file_name"`
	Prompt         string `json:"prompt"`
	MaskFilePath   string `json:"mask_file_path"`
	MaskFileName   string `json:"mask_file_name"`
	Model          string `json:"model"`
	N              int    `json:"n"`
	ResponseFormat string `json:"response_format"`
	Size           string `json:"size"`
	User           string `json:"user"`
}

// CreateEditImageRequest结构体添加方法
func (c *CreateEditImageRequest) GetImageFile() string {
	return strings.TrimRight(c.ImageFilePath, "/") + "/" + c.ImageFileName
}

func (c *CreateEditImageRequest) GetMaskFile() string {
	return strings.TrimRight(c.MaskFilePath, "/") + "/" + c.MaskFileName
}

// CreateEditImageRequest构造函数
func NewCreateEditImageRequest() *CreateEditImageRequest {
	return &CreateEditImageRequest{}
}
