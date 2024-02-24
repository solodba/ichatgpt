package impl

import (
	openapi "github.com/sashabaranov/go-openai"
	"github.com/solodba/ichatgpt/apps/audio"
)

func (i *impl) Req2SpeechConvert(req *audio.CreateSpeechRequest) *openapi.CreateSpeechRequest {
	return &openapi.CreateSpeechRequest{
		Model:          openapi.SpeechModel(req.Model),
		Input:          req.Input,
		Voice:          openapi.SpeechVoice(req.Voice),
		ResponseFormat: openapi.SpeechResponseFormat(req.ResponseFormat),
		Speed:          req.Speed,
	}
}
