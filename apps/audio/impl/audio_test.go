package impl_test

import (
	"context"
	"testing"

	"github.com/solodba/ichatgpt/apps/audio"
	"github.com/solodba/ichatgpt/test/tools"
)

var (
	ctx = context.Background()
)

func TestCreateSpeech(t *testing.T) {
	req := audio.NewCreateSpeechRequest()
	req.Model = "tts-1-hd"
	req.Input = "大家好，我叫沃尔夫冈，来自德国。你今天要去哪里?"
	req.Voice = "alloy"
	req.FilePath = "audio"
	req.FileName = "test.mp3"
	speechResp, err := svc.CreateSpeech(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(speechResp))
}

func TestCreateTranscription(t *testing.T) {
	req := audio.NewCreateTranscriptionRequest()
	req.FileName = "test.mp3"
	req.FilePath = "audio"
	req.Model = "whisper-1"
	req.Language = "zh"
	req.ResponseFormat = "json"
	audioResp, err := svc.CreateTranscription(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(audioResp))
}

func TestCreateTranslation(t *testing.T) {
	req := audio.NewCreateTranslationRequest()
	req.FileName = "test.mp3"
	req.FilePath = "audio"
	req.Model = "whisper-1"
	audioResp, err := svc.CreateTranslation(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(audioResp))
}
