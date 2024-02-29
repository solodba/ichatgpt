package impl

import (
	"context"

	"github.com/solodba/ichatgpt/apps/image"
)

func (i *impl) CreateImage(ctx context.Context, req *image.CreateImageRequest) (*image.CreateImageResponse, error) {
	openaiImageReq := i.Req2ImageConvert(req)
	openaiImageResp, err := i.client.CreateImage(ctx, *openaiImageReq)
	if err != nil {
		return nil, err
	}
	imageResp := i.Resp2ImageConvert(openaiImageResp)
	return imageResp, nil
}
