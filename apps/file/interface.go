package file

import (
	"context"
	"strings"
)

// 模块名称
const (
	AppName = "file"
)

// File功能接口
type Service interface {
	UploadFile(context.Context, *UploadFileRequest) (*FileResponseItem, error)
}

// UploadFileRequest结构体
type UploadFileRequest struct {
	FilePath string `json:"file_path"`
	FileName string `json:"file_name"`
	Purpose  string `json:"purpose"`
}

// UploadFileRequest结构体构造函数
func NewUploadFileRequest() *UploadFileRequest {
	return &UploadFileRequest{}
}

// UploadFileRequest添加方法
func (u *UploadFileRequest) GetFile() string {
	return strings.TrimRight(u.FilePath, "/") + "/" + u.FileName
}
