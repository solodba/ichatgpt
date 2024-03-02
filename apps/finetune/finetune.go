package finetune

// FineTuneJobItem结构体
type FineTuneJobItem struct {
	Id              string           `json:"id"`
	CreatedAt       int64            `json:"created_at"`
	Error           *Error           `json:"error"`
	FineTunedModel  string           `json:"fine_tuned_model"`
	FinishedAt      int64            `json:"finished_at"`
	Hyperparameters *Hyperparameters `json:"hyperparameters"`
	Model           string           `json:"model"`
	Object          string           `json:"object"`
	OrganizationId  string           `json:"organization_id"`
	ResultFiles     []string         `json:"result_files"`
	Status          string           `json:"status"`
	TrainedTokens   int              `json:"trained_tokens"`
	TrainingFile    string           `json:"training_file"`
	ValidationFile  string           `json:"validation_file"`
}

// Error结构体
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Param   string `json:"param"`
}

// Error结构体构造函数
func NewError() *Error {
	return &Error{}
}

// FineTuneJobItem结构体构造函数
func NewFineTuneJobItem() *FineTuneJobItem {
	return &FineTuneJobItem{
		Error:           NewError(),
		Hyperparameters: NewHyperparameters(),
	}
}

// FineTuneJobResponse结构体
type FineTuneJobResponse struct {
	Total           int                `json:"total"`
	FineTuneJobList []*FineTuneJobItem `json:"fine_tune_job_list"`
}

// FineTuneJobResponse结构体构造函数
func NewFineTuneJobResponse() *FineTuneJobResponse {
	return &FineTuneJobResponse{
		FineTuneJobList: make([]*FineTuneJobItem, 0),
	}
}

// FineTuneJobResponse结构体添加方法
func (f *FineTuneJobResponse) AddItems(items ...*FineTuneJobItem) {
	f.FineTuneJobList = append(f.FineTuneJobList, items...)
}
