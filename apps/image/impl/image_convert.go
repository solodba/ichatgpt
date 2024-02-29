package impl

import (
	openapi "github.com/sashabaranov/go-openai"
	"github.com/solodba/ichatgpt/apps/image"
)

// CreateImageRequest转换函数
func (i *impl) Req2ImageConvert(req *image.CreateImageRequest) *openapi.ImageRequest {
	return &openapi.ImageRequest{
		Prompt:         req.Prompt,
		Model:          req.Model,
		N:              req.N,
		Quality:        req.Quality,
		Size:           req.Size,
		Style:          req.Style,
		ResponseFormat: req.ResponseFormat,
		User:           req.User,
	}
}

// CreateImageResponse转换函数
func (i *impl) Resp2ImageConvert(resp openapi.ImageResponse) *image.CreateImageResponse {
	return &image.CreateImageResponse{
		Data: i.Resp2DataInnerConvert(resp.Data),
	}
}

// ImageResponseDataInner转换函数
func (i *impl) Resp2DataInnerConvert(data []openapi.ImageResponseDataInner) []*image.CreateImageResponseItem {
	imageResp := make([]*image.CreateImageResponseItem, 0)
	for _, item := range data {
		imageResp = append(imageResp, &image.CreateImageResponseItem{
			B64Json:       item.B64JSON,
			Url:           item.URL,
			RevisedPrompt: item.RevisedPrompt,
		})
	}
	return imageResp
}
