package impl

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/solodba/ichatgpt/apps/audio"
)

func (i *impl) CreateSpeech(ctx context.Context, req *audio.CreateSpeechRequest) (*audio.CreateSpeechResponse, error) {
	openaiSpeechReq := i.Req2SpeechConvert(req)
	openaiSpeechResp, err := i.client.CreateSpeech(ctx, *openaiSpeechReq)
	if err != nil {
		return nil, err
	}
	speechResp := audio.NewCreateSpeechResponse()
	f, err := os.Create(req.GetFile())
	if err != nil {
		speechResp.Message = fmt.Sprintf("创建[%s]语音文件失败!原因: %s", req.GetFile(), err.Error())
		return speechResp, err
	}
	_, err = io.Copy(f, openaiSpeechResp)
	if err != nil {
		speechResp.Message = fmt.Sprintf("创建[%s]语音文件失败!原因: %s", req.GetFile(), err.Error())
		return speechResp, err
	}
	speechResp.Message = fmt.Sprintf("创建[%s]语音文件成功!", req.GetFile())
	return speechResp, nil
}

func (i *impl) CreateTranscription(ctx context.Context, req *audio.CreateTranscriptionRequest) (*audio.CreateSpeechResponse, error) {
	return nil, nil
}
