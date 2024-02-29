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
	imageReq.Prompt = "A cute baby sea otter"
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
