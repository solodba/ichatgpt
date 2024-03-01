package impl

import (
	"context"

	"github.com/solodba/ichatgpt/apps/model"
)

func (i *impl) ListModels(ctx context.Context, req *model.ListModelsRequest) (*model.ModelsResponse, error) {
	openaiModelsResp, err := i.client.ListModels(ctx)
	if err != nil {
		return nil, err
	}
	modelsResp := i.Resp2ModelConvert(&openaiModelsResp)
	modelsResp.Total = len(modelsResp.ListModels)
	return modelsResp, nil
}

func (i *impl) RetrieveModel(ctx context.Context, req *model.RetrieveModelRequest) (*model.ModelsItem, error) {
	openaiModelsResp, err := i.client.GetModel(ctx, req.Model)
	if err != nil {
		return nil, err
	}
	return &model.ModelsItem{
		Id:      openaiModelsResp.ID,
		Created: openaiModelsResp.CreatedAt,
		Object:  openaiModelsResp.Object,
		OwnerBy: openaiModelsResp.OwnedBy,
	}, nil
}

func (i *impl) DeleteFineTunedModel(ctx context.Context, req *model.DeleteFineTunedModelRequest) (*model.DeleteFineTunedModelResponse, error) {
	openapiResp, err := i.client.DeleteFineTuneModel(ctx, req.Model)
	if err != nil {
		return nil, err
	}
	deleteFineTunedResp := model.NewDeleteFineTunedModelResponse()
	deleteFineTunedResp.Id = openapiResp.ID
	deleteFineTunedResp.Object = openapiResp.Object
	deleteFineTunedResp.Deleted = openapiResp.Deleted
	return deleteFineTunedResp, nil
}
