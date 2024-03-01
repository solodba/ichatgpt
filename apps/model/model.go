package model

// ListModelsItem结构体
type ListModelsItem struct {
	Id      string `json:"id"`
	Created int64  `json:"created"`
	Object  string `json:"object"`
	OwnerBy string `json:"owned_by"`
}

// ListModelsResponse结构体
type ListModelsResponse struct {
	Total      int               `json:"total"`
	ListModels []*ListModelsItem `json:"list_models"`
}

// ListModelsResponse结构体初始化函数
func NewListModelsResponse() *ListModelsResponse {
	return &ListModelsResponse{
		ListModels: make([]*ListModelsItem, 0),
	}
}

// ListModelsResponse添加方法
func (m *ListModelsResponse) AddItems(items ...*ListModelsItem) {
	m.ListModels = append(m.ListModels, items...)
}
