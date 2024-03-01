package impl

import (
	"context"

	"github.com/solodba/ichatgpt/apps/model"
)

func (i *impl) ListModels(ctx context.Context, req *model.ListModelsRequest) (*model.ListModelsResponse, error) {
	openaiModelsResp, err := i.client.ListModels(ctx)
	if err != nil {
		return nil, err
	}
	listModelsResp := i.Resp2ModelConvert(&openaiModelsResp)
	listModelsResp.Total = len(listModelsResp.ListModels)
	return listModelsResp, nil
}
