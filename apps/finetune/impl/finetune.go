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

func (i *impl) ListFineTuneJobEvents(ctx context.Context, req *finetune.ListFineTuneJobEventsRequest) (*finetune.FineTuneJobEventResponse, error) {
	openaiFineTuneJobEvents, err := i.client.ListFineTuningJobEvents(ctx, req.FineTuningJobId)
	if err != nil {
		return nil, err
	}
	fineTuneJobEventsResp := i.OpenaiFineTuneJobEvents2Resp(openaiFineTuneJobEvents)
	fineTuneJobEventsResp.Total = len(fineTuneJobEventsResp.FineTuneJobEventList)
	return fineTuneJobEventsResp, nil
}

func (i *impl) RetrieveFineTuneJob(ctx context.Context, req *finetune.RetrieveFineTuneJobRequest) (*finetune.FineTuneJobItem, error) {
	openaiFineTuneJobResp, err := i.client.RetrieveFineTuningJob(ctx, req.FineTuningJobId)
	if err != nil {
		return nil, err
	}
	return i.OpenaiFineTuneJob2Resp(&openaiFineTuneJobResp), nil
}
