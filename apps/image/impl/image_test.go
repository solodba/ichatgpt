package impl_test

import (
	"context"
	"testing"

	"github.com/solodba/ichatgpt/apps/image"
	"github.com/solodba/ichatgpt/test/tools"
)

var (
	ctx = context.Background()
)

func TestCreateImage(t *testing.T) {
	imageReq := image.NewCreateImageRequest()
	imageReq.Prompt = "贝雷帽"
	imageReq.Model = "dall-e-3"
	imageReq.N = 1
	imageReq.Size = "1024x1024"
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
	imageEditReq.ImageFileName = "小狗.png"
	imageEditReq.MaskFilePath = "image"
	imageEditReq.MaskFileName = "贝雷帽.png"
	imageEditReq.Prompt = "小狗带贝雷帽"
	imageEditReq.N = 2
	imageEditReq.Size = "1024x1024"
	imageEditReq.ResponseFormat = "url"
	imageEditResp, err := svc.CreateEditImage(ctx, imageEditReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(imageEditResp))
	t.Log(imageEditResp.Data[0].Url)
}
