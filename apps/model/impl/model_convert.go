package impl

import (
	openapi "github.com/sashabaranov/go-openai"
	"github.com/solodba/ichatgpt/apps/model"
)

func (i *impl) Resp2ModelConvert(req *openapi.ModelsList) *model.ListModelsResponse {
	return &model.ListModelsResponse{
		ListModels: i.Resp2Model(req.Models),
	}
}

func (i *impl) Resp2Model(openaiModel []openapi.Model) []*model.ListModelsItem {
	listModels := make([]*model.ListModelsItem, 0)
	for _, item := range openaiModel {
		listModels = append(listModels, &model.ListModelsItem{
			Id:      item.ID,
			Created: item.CreatedAt,
			Object:  item.Object,
			OwnerBy: item.OwnedBy,
		})
	}
	return listModels
}
