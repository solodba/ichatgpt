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

func (i *impl) Req2TranscriptionConvert(req *audio.CreateTranscriptionRequest) *openapi.AudioRequest {
	return &openapi.AudioRequest{
		Model:       req.Model,
		FilePath:    req.GetFile(),
		Prompt:      req.Prompt,
		Temperature: req.Temperature,
		Language:    req.Language,
		Format:      openapi.AudioResponseFormat(req.ResponseFormat),
	}
}

func (i *impl) Resp2AudioConvert(audioResp *openapi.AudioResponse) *audio.CreateAudioResponse {
	return &audio.CreateAudioResponse{
		Language: audioResp.Language,
		Duration: audioResp.Duration,
		Words:    []*audio.WordItem{},
		Segments: i.Resp2AudioSegmentsConvert(audioResp),
		Text:     audioResp.Text,
	}
}

func (i *impl) Resp2AudioSegmentsConvert(audioRespSegments *openapi.AudioResponse) []*audio.SegmentItem {
	audioSegments := make([]*audio.SegmentItem, 0)
	for _, item := range audioRespSegments.Segments {
		audioSegments = append(audioSegments, &audio.SegmentItem{
			Id:               item.ID,
			Seek:             item.Seek,
			Start:            item.Start,
			End:              item.End,
			Text:             item.Text,
			Tokens:           item.Tokens,
			Temperature:      item.Temperature,
			AvgLogprob:       item.AvgLogprob,
			CompressionRatio: item.CompressionRatio,
			NoSpeechProb:     item.NoSpeechProb,
		})
	}
	return audioSegments
}

func (i *impl) Req2TranslationConvert(req *audio.CreateTranslationRequest) *openapi.AudioRequest {
	return &openapi.AudioRequest{
		Model:       req.Model,
		FilePath:    req.GetFile(),
		Prompt:      req.Prompt,
		Temperature: req.Temperature,
		Format:      openapi.AudioResponseFormat(req.ResponseFormat),
	}
}
