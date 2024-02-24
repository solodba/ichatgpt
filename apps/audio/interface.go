package audio

import (
	"context"
	"strings"
)

// 模块名称
const (
	AppName = "audio"
)

// Audio功能接口
type Service interface {
	CreateSpeech(context.Context, *CreateSpeechRequest) (*CreateSpeechResponse, error)
}

// CreateSpeechRequest结构体
type CreateSpeechRequest struct {
	Model          string  `json:"model"`
	Input          string  `json:"input"`
	Voice          string  `json:"voice"`
	ResponseFormat string  `json:"response_format"`
	Speed          float64 `json:"speed"`
	FilePath       string  `json:"file_path"`
	FileName       string  `json:"file_name"`
}

// CreateSpeechRequest结构体构造函数
func NewCreateSpeechRequest() *CreateSpeechRequest {
	return &CreateSpeechRequest{}
}

// 获取文件路径
func (c *CreateSpeechRequest) GetFile() string {
	return strings.TrimRight(c.FilePath, "/") + "/" + c.FileName
}
