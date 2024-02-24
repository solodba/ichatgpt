package impl_test

import (
	"context"
	"os"
	"testing"

	"github.com/solodba/ichatgpt/apps/audio"
	"github.com/solodba/ichatgpt/test/tools"
)

var (
	ctx = context.Background()
)

func TestCreateSpeech(t *testing.T) {
	err := os.Setenv("http_proxy", "http://192.168.1.3:7890")
	if err != nil {
		t.Fatal(err)
	}
	err = os.Setenv("https_proxy", "http://192.168.1.3:7890")
	if err != nil {
		t.Fatal(err)
	}
	req := audio.NewCreateSpeechRequest()
	req.Model = "tts-1-hd"
	req.Input = "The quick brown fox jumped over the lazy dog."
	req.Voice = "alloy"
	req.FilePath = "audio"
	req.FileName = "test.mp3"
	speechResp, err := svc.CreateSpeech(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(speechResp))
}
