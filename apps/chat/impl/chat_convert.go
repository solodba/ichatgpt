package impl

import (
	openapi "github.com/sashabaranov/go-openai"
	"github.com/solodba/ichatgpt/apps/chat"
)

func (i *impl) Req2ChatConvert(req *chat.CreateChatRequest) *openapi.ChatCompletionRequest {
	return &openapi.ChatCompletionRequest{
		Model:    req.Model,
		Messages: i.Req2ChatMessagesConvert(req.Messages),
	}
}

func (i *impl) Req2ChatMessagesConvert(messages []*chat.MessagesItem) []openapi.ChatCompletionMessage {
	chatCompletionMessage := make([]openapi.ChatCompletionMessage, 0)
	for _, item := range messages {
		chatCompletionMessage = append(chatCompletionMessage, openapi.ChatCompletionMessage{
			Role:    item.Role,
			Content: item.Content,
		})
	}
	return chatCompletionMessage
}

func (i *impl) Resp2ChatConvert(resp openapi.ChatCompletionResponse) *chat.CreateChatResponse {
	return &chat.CreateChatResponse{
		Id:                resp.ID,
		Object:            resp.Object,
		Created:           resp.Created,
		Model:             resp.Model,
		SystemFingerprint: resp.SystemFingerprint,
		Choices:           i.Resp2ChatChoicesConvert(resp.Choices),
		Usage:             i.Resp2ChatUsageConvert(resp.Usage),
	}
}

func (i *impl) Resp2ChatChoicesConvert(choices []openapi.ChatCompletionChoice) []*chat.ChoicesItem {
	chatChoices := make([]*chat.ChoicesItem, 0)
	for _, item := range choices {
		chatChoices = append(chatChoices, &chat.ChoicesItem{
			FinishReason: string(item.FinishReason),
			Index:        item.Index,
			Messages:     i.Resp2ChatChoicesMessagesConvert(item.Message),
			LogProbs:     i.Resp2ChatChoicesLogProbsConvert(item.LogProbs),
		})
	}
	return chatChoices
}

func (i *impl) Resp2ChatChoicesMessagesConvert(message openapi.ChatCompletionMessage) *chat.MessagesItem {
	return &chat.MessagesItem{
		Role:    message.Role,
		Content: message.Content,
	}
}

func (i *impl) Resp2ChatChoicesLogProbsConvert(logprobs *openapi.LogProbs) *chat.LogProbs {
	if logprobs == nil {
		return nil
	}
	return &chat.LogProbs{
		Content: i.Resp2ChatChoicesLogProbsContentConvert(logprobs.Content),
	}
}

func (i *impl) Resp2ChatChoicesLogProbsContentConvert(content []openapi.LogProb) []*chat.ContentItem {
	chatContent := make([]*chat.ContentItem, 0)
	for _, item := range content {
		chatContent = append(chatContent, &chat.ContentItem{
			Token:       item.Token,
			Logprob:     item.LogProb,
			Bytes:       item.Bytes,
			TopLogprobs: i.Resp2ChatChoicesTopLogProbs(item.TopLogProbs),
		})
	}
	return chatContent
}

func (i *impl) Resp2ChatChoicesTopLogProbs(topLogProbs []openapi.TopLogProbs) []*chat.TopLogprobsItem {
	chatTopLogProbs := make([]*chat.TopLogprobsItem, 0)
	for _, item := range topLogProbs {
		chatTopLogProbs = append(chatTopLogProbs, &chat.TopLogprobsItem{
			Token:   item.Token,
			Logprob: item.LogProb,
			Bytes:   item.Bytes,
		})
	}
	return chatTopLogProbs
}

func (i *impl) Resp2ChatUsageConvert(usage openapi.Usage) *chat.Usage {
	return &chat.Usage{
		CompletionTokens: usage.CompletionTokens,
		PromptTokens:     usage.PromptTokens,
		TotalTokens:      usage.TotalTokens,
	}
}
