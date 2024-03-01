package model

import "context"

// 模块名称
const (
	AppName = "model"
)

// Model功能接口
type Service interface {
	ListModels(context.Context, *ListModelsRequest) (*ModelsResponse, error)
	RetrieveModel(context.Context, *RetrieveModelRequest) (*ModelsItem, error)
}

// ListModelsRequest结构体
type ListModelsRequest struct {
}

// ListModelsRequest结构体初始化函数
func NewListModelsRequest() *ListModelsRequest {
	return &ListModelsRequest{}
}

// RetrieveModelRequest结构体
type RetrieveModelRequest struct {
	Model string `json:"model"`
}

// RetrieveModelRequest结构体初始化函数
func NewRetrieveModelRequest() *RetrieveModelRequest {
	return &RetrieveModelRequest{}
}
