package impl

import (
	"context"

	"github.com/solodba/ichatgpt/apps/chat"
)

func (i *impl) CreateChat(ctx context.Context, req *chat.CreateChatRequest) (*chat.CreateChatResponse, error) {
	openaiChatReq := i.Req2ChatConvert(req)
	openaiChatResp, err := i.client.CreateChatCompletion(ctx, *openaiChatReq)
	if err != nil {
		return nil, err
	}
	return i.Resp2ChatConvert(openaiChatResp), nil
}
