package impl

import (
	openapi "github.com/sashabaranov/go-openai"
	"github.com/solodba/ichatgpt/apps/file"
)

func (i *impl) Req2UploadFileConvert(req *file.UploadFileRequest) *openapi.FileRequest {
	return &openapi.FileRequest{
		FilePath: req.GetFile(),
		Purpose:  req.Purpose,
	}
}

func (i *impl) Resp2UploadFileConvert(resp *openapi.File) *file.FileResponseItem {
	return &file.FileResponseItem{
		Id:            resp.ID,
		Bytes:         resp.Bytes,
		CreateAt:      resp.CreatedAt,
		Filename:      resp.FileName,
		Object:        resp.Object,
		Purpose:       resp.Purpose,
		Status:        resp.Status,
		StatusDetails: resp.StatusDetails,
	}
}

func (i *impl) Resp2ListFilesConvert(resp openapi.FilesList) *file.FileResponse {
	fileResp := file.NewFileResponse()
	for _, item := range resp.Files {
		fileResp.AddItems(i.Resp2UploadFileConvert(&item))
	}
	return fileResp
}
