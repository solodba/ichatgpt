package impl_test

import (
	"testing"

	"github.com/solodba/ichatgpt/apps/chat"
	"github.com/solodba/ichatgpt/test/tools"
)

func TestCreateChat(t *testing.T) {
	chatReq := chat.NewCreateChatRequest()
	chatReq.Model = "gpt-3.5-turbo"
	item1 := chat.NewMessagesItem()
	item1.Role = "system"
	item1.Content = "You are a helpful assistant."
	item2 := chat.NewMessagesItem()
	item2.Role = "user"
	item2.Content = "用中文介绍一下自己"
	chatReq.AddItems(item1, item2)
	chatResp, err := svc.CreateChat(ctx, chatReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(chatResp))
}
