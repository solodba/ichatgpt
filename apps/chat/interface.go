package chat

import "context"

// 模块名称
const (
	AppName = "chat"
)

// Chat功能接口
type Service interface {
	CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error)
}

// CreateChatRequest结构体
type CreateChatRequest struct {
	Model    string          `json:"model"`
	Messages []*MessagesItem `json:"messages"`
}

// MessagesItem结构体
type MessagesItem struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// MessagesItem结构体构造函数
func NewMessagesItem() *MessagesItem {
	return &MessagesItem{}
}

// CreateChatRequest结构体构造函数
func NewCreateChatRequest() *CreateChatRequest {
	return &CreateChatRequest{
		Messages: make([]*MessagesItem, 0),
	}
}

// CreateChatRequest结构体添加方法
func (c *CreateChatRequest) AddItems(items ...*MessagesItem) {
	c.Messages = append(c.Messages, items...)
}
