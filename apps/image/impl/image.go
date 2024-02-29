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

func (i *impl) CreateEditImage(ctx context.Context, req *image.CreateEditImageRequest) (*image.CreateImageResponse, error) {
	openaiImageEditReq, err := i.Req2ImageEditConvert(req)
	if err != nil {
		return nil, err
	}
	openaiImageEditResp, err := i.client.CreateEditImage(ctx, *openaiImageEditReq)
	if err != nil {
		return nil, err
	}
	imageEditResp := i.Resp2ImageConvert(openaiImageEditResp)
	return imageEditResp, nil
}
