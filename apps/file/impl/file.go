package impl

import (
	"context"

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
