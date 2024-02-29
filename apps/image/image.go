package image

// CreateImageResponseItem结构体
type CreateImageResponseItem struct {
	B64Json       string `json:"b64_json"`
	Url           string `json:"url"`
	RevisedPrompt string `json:"revised_prompt"`
}

// CreateImageResponse结构体
type CreateImageResponse struct {
	Data []*CreateImageResponseItem `json:"data"`
}

// CreateImageResponse结构体构造函数
func NewCreateImageResponseItem() *CreateImageResponseItem {
	return &CreateImageResponseItem{}
}

// CreateImageResponse结构体构造函数
func NewCreateImageResponse() *CreateImageResponse {
	return &CreateImageResponse{
		Data: make([]*CreateImageResponseItem, 0),
	}
}

// CreateImageResponse添加方法
func (c *CreateImageResponse) AddItems(item ...*CreateImageResponseItem) {
	c.Data = append(c.Data, item...)
}
