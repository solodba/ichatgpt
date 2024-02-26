package audio

// CreateSpeechResponse结构体
type CreateSpeechResponse struct {
	Message string `json:"message"`
}

// CreateSpeechResponse结构体构造函数
func NewCreateSpeechResponse() *CreateSpeechResponse {
	return &CreateSpeechResponse{}
}

// CreateTranscriptionResponse结构体
type CreateTranscriptionResponse struct {
	Language string         `json:"language"`
	Duration float64        `json:"duration"`
	Text     string         `json:"text"`
	Words    []*WordItem    `json:"words"`
	Segments []*SegmentItem `json:"segments"`
}

// WordItem结构体
type WordItem struct {
	Word  string `json:"word"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

// SegmentItem结构体
type SegmentItem struct {
	Id               int    `json:"id"`
	Seek             int    `json:"seek"`
	Start            int    `json:"start"`
	End              int    `json:"end"`
	Text             string `json:"text"`
	Tokens           []int  `json:"tokens"`
	Temperature      int    `json:"temperature"`
	AvgLogprob       int    `json:"avg_logprob"`
	CompressionRatio int    `json:"compression_ratio"`
	NoSpeechProb     int    `json:"no_speech_prob"`
}

// WordItem结构体构造函数
func NewWordItem() *WordItem {
	return &WordItem{}
}

// SegmentItem结构体构造函数
func NewSegmentItem() *SegmentItem {
	return &SegmentItem{
		Tokens: make([]int, 0),
	}
}

// CreateTranscriptionResponse结构体构造函数
func NewCreateTranscriptionResponse() *CreateTranscriptionResponse {
	return &CreateTranscriptionResponse{
		Words:    make([]*WordItem, 0),
		Segments: make([]*SegmentItem, 0),
	}
}
