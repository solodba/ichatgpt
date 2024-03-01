package impl

import (
	"context"
	"fmt"

	"github.com/solodba/ichatgpt/apps/file"
)

func (i *impl) UploadFile(ctx context.Context, req *file.UploadFileRequest) (*file.FileResponseItem, error) {
	uploadFileReq := i.Req2UploadFileConvert(req)
	openaiFile, err := i.client.CreateFile(ctx, *uploadFileReq)
	if err != nil {
		return nil, err
	}
	return i.Resp2UploadFileConvert(&openaiFile), nil
}

func (i *impl) ListFile(ctx context.Context, req *file.ListFileRequest) (*file.FileResponse, error) {
	openaiListFiles, err := i.client.ListFiles(ctx)
	if err != nil {
		return nil, err
	}
	openaiFileResp := i.Resp2ListFilesConvert(openaiListFiles)
	openaiFileResp.Total = len(openaiFileResp.Files)
	return openaiFileResp, nil
}

func (i *impl) RetrieveFile(ctx context.Context, req *file.RetrieveFileRequest) (*file.FileResponseItem, error) {
	openaiRetrieveFile, err := i.client.GetFile(ctx, req.FileId)
	if err != nil {
		return nil, err
	}
	return i.Resp2UploadFileConvert(&openaiRetrieveFile), nil
}

func (i *impl) DeleteFile(ctx context.Context, req *file.DeleteFileRequest) (*file.DeleteFileResponse, error) {
	err := i.client.DeleteFile(ctx, req.FileId)
	if err != nil {
		return nil, err
	}
	deleteFileResp := file.NewDeleteFileResponse()
	deleteFileResp.Message = fmt.Sprintf("删除文件[%s]成功!", req.FileId)
	return deleteFileResp, nil
}
