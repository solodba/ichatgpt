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
