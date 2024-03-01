package model

// ModelsItem结构体
type ModelsItem struct {
	Id      string `json:"id"`
	Created int64  `json:"created"`
	Object  string `json:"object"`
	OwnerBy string `json:"owned_by"`
}

// ModelsResponse结构体
type ModelsResponse struct {
	Total      int           `json:"total"`
	ListModels []*ModelsItem `json:"list_models"`
}

// ModelsResponse结构体初始化函数
func NewModelsResponse() *ModelsResponse {
	return &ModelsResponse{
		ListModels: make([]*ModelsItem, 0),
	}
}

// ModelsResponse添加方法
func (m *ModelsResponse) AddItems(items ...*ModelsItem) {
	m.ListModels = append(m.ListModels, items...)
}

// DeleteFineTunedModelResponse结构体
type DeleteFineTunedModelResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

// DeleteFineTunedModelResponse结构体构造函数
func NewDeleteFineTunedModelResponse() *DeleteFineTunedModelResponse {
	return &DeleteFineTunedModelResponse{}
}
