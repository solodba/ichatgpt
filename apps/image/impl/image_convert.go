package impl

import (
	"fmt"
	"os"
	"path/filepath"

	openapi "github.com/sashabaranov/go-openai"
	"github.com/solodba/ichatgpt/apps/image"

	myimage "image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
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
		Created: resp.Created,
		Data:    i.Resp2DataInnerConvert(resp.Data),
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

// CreateEditImageRequest转换函数
func (i *impl) Req2ImageEditConvert(req *image.CreateEditImageRequest) (*openapi.ImageEditRequest, error) {
	imagefs, err := os.Open(req.GetImageFile())
	if err != nil {
		return nil, err
	}
	maskfs, err := os.Open(req.GetMaskFile())
	if err != nil {
		return nil, err
	}
	return &openapi.ImageEditRequest{
		Image:          imagefs,
		Mask:           maskfs,
		Prompt:         req.Prompt,
		Model:          req.Model,
		N:              req.N,
		Size:           req.Size,
		ResponseFormat: req.ResponseFormat,
	}, nil
}

func ImageConvert(filePath string, newFilePath string) error {
	image, err := LoadImage(filePath)
	if err != nil {
		return err
	}
	imageRGBA := ConvertToRGBA(image)
	err = SaveImage(imageRGBA, newFilePath)
	if err != nil {
		return err
	}
	return nil
}

func LoadImage(path string) (myimage.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := myimage.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func ConvertToRGBA(img myimage.Image) *myimage.RGBA {
	bounds := img.Bounds()
	rgba := myimage.NewRGBA(bounds)

	draw.Draw(rgba, bounds, &myimage.Uniform{color.White}, myimage.Point{}, draw.Over)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Over)

	return rgba
}

func SaveImage(img myimage.Image, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	ext := filepath.Ext(outputPath)
	switch ext {
	case ".png":
		return png.Encode(file, img)
	case ".jpg", ".jpeg":
		return jpeg.Encode(file, img, nil)
	default:
		return fmt.Errorf("unsupported image format: %s", ext)
	}
}
