package image

import "context"

// 模块名称
const (
	AppName = "image"
)

// Image功能接口
type Service interface {
	CreateImage(context.Context, *CreateImageRequest) (*CreateImageResponse, error)
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
