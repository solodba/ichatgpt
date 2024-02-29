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
	CreateTranscription(context.Context, *CreateTranscriptionRequest) (*CreateTranscriptionResponse, error)
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

// CreateTranscriptionRequest结构体
type CreateTranscriptionRequest struct {
	FileName       string  `json:"file_name"`
	FilePath       string  `json:"file_path"`
	Model          string  `json:"model"`
	Language       string  `json:"language"`
	Prompt         string  `json:"prompt"`
	ResponseFormat string  `json:"response_format"`
	Temperature    float32 `json:"temperature"`
}

// CreateTranscriptionRequest结构体构造函数
func NewCreateTranscriptionRequest() *CreateTranscriptionRequest {
	return &CreateTranscriptionRequest{}
}

// CreateTranscriptionRequest获取文件路径
func (c *CreateTranscriptionRequest) GetFile() string {
	return strings.TrimRight(c.FilePath, "/") + "/" + c.FileName
}
