package impl

import (
	openapi "github.com/sashabaranov/go-openai"
	"github.com/solodba/ichatgpt/apps/finetune"
)

func (i *impl) Req2FineTuneJobConvert(req *finetune.CreateFineTuneJobRequest) *openapi.FineTuningJobRequest {
	return &openapi.FineTuningJobRequest{
		TrainingFile:    req.TrainingFile,
		ValidationFile:  req.ValidationFile,
		Model:           req.Model,
		Hyperparameters: i.Req2HyperparametersConvert(req.Hyperparameters),
		Suffix:          req.Suffix,
	}
}

func (i *impl) Req2HyperparametersConvert(hy *finetune.Hyperparameters) *openapi.Hyperparameters {
	if hy == nil {
		return nil
	}
	return &openapi.Hyperparameters{
		Epochs: hy.NEpochs,
	}
}

func (i *impl) OpenaiFineTuneJob2Resp(resp *openapi.FineTuningJob) *finetune.FineTuneJobItem {
	return &finetune.FineTuneJobItem{
		Id:              resp.ID,
		CreatedAt:       resp.CreatedAt,
		Error:           finetune.NewError(),
		FineTunedModel:  resp.FineTunedModel,
		FinishedAt:      resp.FinishedAt,
		Hyperparameters: i.Resp2HyperparametersConvert(resp.Hyperparameters),
		Model:           resp.Model,
		Object:          resp.Object,
		OrganizationId:  resp.OrganizationID,
		ResultFiles:     resp.ResultFiles,
		Status:          resp.Status,
		TrainedTokens:   resp.TrainedTokens,
		TrainingFile:    resp.TrainingFile,
		ValidationFile:  resp.ValidationFile,
	}
}

func (i *impl) Resp2HyperparametersConvert(hy openapi.Hyperparameters) *finetune.Hyperparameters {
	return &finetune.Hyperparameters{
		NEpochs: hy.Epochs,
	}
}

func (i *impl) OpenaiFineTuneJobEvents2Resp(resp openapi.FineTuningJobEventList) *finetune.FineTuneJobEventResponse {
	fineTuneJobEventList := make([]*finetune.FineTuneJobEventItem, 0)
	for _, item := range resp.Data {
		fineTuneJobEventList = append(fineTuneJobEventList, &finetune.FineTuneJobEventItem{
			CreatedAt: item.CreatedAt,
			Level:     item.Level,
			Message:   item.Message,
			Object:    item.Object,
		})
	}
	return &finetune.FineTuneJobEventResponse{
		FineTuneJobEventList: fineTuneJobEventList,
	}
}
