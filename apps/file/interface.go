package file

import (
	"context"
	"strings"

	"github.com/emicklei/go-restful/v3"
)

// 模块名称
const (
	AppName = "file"
)

// File功能接口
type Service interface {
	UploadFile(context.Context, *UploadFileRequest) (*FileResponseItem, error)
	ListFile(context.Context, *ListFileRequest) (*FileResponse, error)
	RetrieveFile(context.Context, *RetrieveFileRequest) (*FileResponseItem, error)
	DeleteFile(context.Context, *DeleteFileRequest) (*DeleteFileResponse, error)
	RetrieveFileContent(context.Context, *RetrieveFileContentRequest) (*RetrieveFileContentResponse, error)
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

// ListFileRequest结构体
type ListFileRequest struct {
	Purpose string `json:"purpose"`
}

// ListFileRequest结构体结构体
func NewListFileRequest() *ListFileRequest {
	return &ListFileRequest{}
}

// RetrieveFileRequest结构体
type RetrieveFileRequest struct {
	FileId string `json:"file_id"`
}

// RetrieveFileRequest结构体构造函数
func NewRetrieveFileRequest() *RetrieveFileRequest {
	return &RetrieveFileRequest{}
}

// DeleteFileRequest结构体
type DeleteFileRequest struct {
	FileId string `json:"file_id"`
}

// DeleteFileRequest结构体构造函数
func NewDeleteFileRequest() *DeleteFileRequest {
	return &DeleteFileRequest{}
}

// RetrieveFileContentRequest结构体
type RetrieveFileContentRequest struct {
	FileId   string `json:"file_id"`
	FilePath string `json:"file_path"`
	FileName string `json:"file_name"`
}

// RetrieveFileContentRequest结构体构造函数
func NewRetrieveFileContentRequest() *RetrieveFileContentRequest {
	return &RetrieveFileContentRequest{}
}

// RetrieveFileContentRequest添加方法
func (r *RetrieveFileContentRequest) GetFile() string {
	return strings.TrimRight(r.FilePath, "/") + "/" + r.FileName
}

// NewListFileRequestFromRestful结构体
func NewListFileRequestFromRestful(r *restful.Request) *ListFileRequest {
	return &ListFileRequest{
		Purpose: r.QueryParameter("purpose"),
	}
}

// RetrieveFileRequest结构体构造函数
func NewRetrieveFileRequestFromRestful(r *restful.Request) *RetrieveFileRequest {
	return &RetrieveFileRequest{
		FileId: r.PathParameter("file_id"),
	}
}

// RetrieveFileContentRequest结构体构造函数
func NewRetrieveFileContentRequestFromRestful(r *restful.Request) *RetrieveFileContentRequest {
	return &RetrieveFileContentRequest{
		FileId:   r.PathParameter("file_id"),
		FilePath: r.QueryParameter("file_path"),
		FileName: r.QueryParameter("file_name"),
	}
}

// DeleteFileRequest结构体构造函数
func NewDeleteFileRequestFromRestful(r *restful.Request) *DeleteFileRequest {
	return &DeleteFileRequest{
		FileId: r.PathParameter("file_id"),
	}
}
