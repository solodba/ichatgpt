package impl

import (
	"context"

	"github.com/solodba/ichatgpt/apps/finetune"
)

func (i *impl) CreateFineTuneJob(ctx context.Context, req *finetune.CreateFineTuneJobRequest) (*finetune.FineTuneJobItem, error) {
	openaiFineTuneJobReq := i.Req2FineTuneJobConvert(req)
	openaiFineTuneJobResp, err := i.client.CreateFineTuningJob(ctx, *openaiFineTuneJobReq)
	if err != nil {
		return nil, err
	}
	return i.OpenaiFineTuneJob2Resp(&openaiFineTuneJobResp), nil
}
