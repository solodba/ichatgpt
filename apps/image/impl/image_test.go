package impl_test

import (
	"testing"

	"github.com/solodba/ichatgpt/apps/image"
	"github.com/solodba/ichatgpt/test/tools"
)

func TestCreateImage(t *testing.T) {
	imageReq := image.NewCreateImageRequest()
	imageReq.Prompt = "cap"
	imageReq.Model = "dall-e-2"
	imageReq.N = 1
	imageReq.Size = "256x256"
	imageResp, err := svc.CreateImage(ctx, imageReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(imageResp))
	t.Log(imageResp.Data[0].Url)
}

func TestCreateEditImage(t *testing.T) {
	imageEditReq := image.NewCreateEditImageRequest()
	imageEditReq.ImageFilePath = "image"
	imageEditReq.ImageFileName = "dog.rgba"
	imageEditReq.MaskFilePath = "image"
	imageEditReq.MaskFileName = "mask.rgba"
	imageEditReq.Prompt = "小狗带帽子"
	imageEditReq.N = 2
	imageEditReq.Size = "256x256"
	imageEditReq.ResponseFormat = "url"
	imageEditResp, err := svc.CreateEditImage(ctx, imageEditReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(imageEditResp))
	t.Log(imageEditResp.Data[0].Url)
}
