package impl

import (
	openapi "github.com/sashabaranov/go-openai"
	"github.com/solodba/ichatgpt/apps/model"
)

func (i *impl) Resp2ModelConvert(req *openapi.ModelsList) *model.ModelsResponse {
	return &model.ModelsResponse{
		ListModels: i.Resp2Model(req.Models),
	}
}

func (i *impl) Resp2Model(openaiModel []openapi.Model) []*model.ModelsItem {
	models := make([]*model.ModelsItem, 0)
	for _, item := range openaiModel {
		models = append(models, &model.ModelsItem{
			Id:      item.ID,
			Created: item.CreatedAt,
			Object:  item.Object,
			OwnerBy: item.OwnedBy,
		})
	}
	return models
}
