package finetune

import "context"

// 模块名称
const (
	AppName = "finetune"
)

// Model功能接口
type Service interface {
	CreateFineTuneJob(context.Context, *CreateFineTuneJobRequest) (*FineTuneJobItem, error)
}

// CreateFineTuneJobRequest结构体
type CreateFineTuneJobRequest struct {
	Model           string           `json:"model"`
	TrainingFile    string           `json:"training_file"`
	Hyperparameters *Hyperparameters `json:"hyperparameters"`
	Suffix          string           `json:"suffix"`
	ValidationFile  string           `json:"validation_file"`
}

// Hyperparameters结构体
type Hyperparameters struct {
	BatchSize              any `json:"batch_size"`
	LearningRateMultiplier any `json:"learning_rate_multiplier"`
	NEpochs                any `json:"n_epochs"`
}

// CreateFineTuneJobRequest结构体构造函数
func NewCreateFineTuneJobRequest() *CreateFineTuneJobRequest {
	return &CreateFineTuneJobRequest{
		Hyperparameters: NewHyperparameters(),
	}
}

// Hyperparameters结构体构造函数
func NewHyperparameters() *Hyperparameters {
	return &Hyperparameters{}
}
