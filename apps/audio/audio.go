package audio

// CreateSpeechResponse结构体
type CreateSpeechResponse struct {
	Message string `json:"message"`
}

// CreateSpeechResponse结构体构造函数
func NewCreateSpeechResponse() *CreateSpeechResponse {
	return &CreateSpeechResponse{}
}
