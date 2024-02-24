package conf

import (
	openapi "github.com/sashabaranov/go-openai"
)

func (c *ChatGPT) GetClient() *openapi.Client {
	return openapi.NewClient(c.OpenaiApiKey)
}
