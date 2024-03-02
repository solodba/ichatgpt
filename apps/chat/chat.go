package chat

// CreateChatResponse结构体
type CreateChatResponse struct {
	Id                string         `json:"id"`
	Choices           []*ChoicesItem `json:"choices"`
	Object            string         `json:"object"`
	Created           int64          `json:"created"`
	Model             string         `json:"model"`
	SystemFingerprint string         `json:"system_fingerprint"`
	Usage             *Usage         `json:"usage"`
}

// ChoicesItem结构体
type ChoicesItem struct {
	FinishReason string        `json:"finish_reason"`
	Index        int           `json:"index"`
	Messages     *MessagesItem `json:"messages"`
	LogProbs     *LogProbs     `json:"logprobs"`
}

// ChoicesItem结构体构造函数
func NewChoicesItem() *ChoicesItem {
	return &ChoicesItem{
		Messages: NewMessagesItem(),
		LogProbs: NewLogProbs(),
	}
}

// CreateChatResponse结构体构造函数
func NewCreateChatResponse() *CreateChatResponse {
	return &CreateChatResponse{
		Choices: make([]*ChoicesItem, 0),
		Usage:   NewUsage(),
	}
}

// LogProbs结构体
type LogProbs struct {
	Content []*ContentItem `json:"content"`
}

// ContentItem结构体
type ContentItem struct {
	Token       string             `json:"token"`
	Logprob     float64            `json:"logprob"`
	Bytes       []byte             `json:"bytes"`
	TopLogprobs []*TopLogprobsItem `json:"top_logprobs"`
}

// TopLogprobsItem结构体
type TopLogprobsItem struct {
	Token   string  `json:"token"`
	Logprob float64 `json:"logprob"`
	Bytes   []byte  `json:"bytes"`
}

// Usage结构体
type Usage struct {
	CompletionTokens int `json:"completion_tokens"`
	PromptTokens     int `json:"prompt_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// LogProbs结构体构造函数
func NewLogProbs() *LogProbs {
	return &LogProbs{
		Content: make([]*ContentItem, 0),
	}
}

// ContentItem结构体构造函数
func NewContentItem() *ContentItem {
	return &ContentItem{
		Bytes:       make([]byte, 0),
		TopLogprobs: make([]*TopLogprobsItem, 0),
	}
}

// TopLogprobsItem结构体
func NewTopLogprobsItem() *TopLogprobsItem {
	return &TopLogprobsItem{
		Bytes: make([]byte, 0),
	}
}

// Usage结构体构造函数
func NewUsage() *Usage {
	return &Usage{}
}
