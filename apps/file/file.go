package file

// FileResponse结构体
type FileResponse struct {
	Total int                 `json:"total"`
	Files []*FileResponseItem `json:"files"`
}

// FileResponseItem结构体
type FileResponseItem struct {
	Id            string `json:"id"`
	Bytes         int    `json:"bytes"`
	CreateAt      int64  `json:"created_at"`
	Filename      string `json:"filename"`
	Object        string `json:"object"`
	Purpose       string `json:"purpose"`
	Status        string `json:"status"`
	StatusDetails string `json:"status_details"`
}

// FileResponse结构体构造函数
func NewFileResponse() *FileResponse {
	return &FileResponse{
		Files: make([]*FileResponseItem, 0),
	}
}

// FileResponseItem结构体构造函数
func NewFileResponseItem() *FileResponseItem {
	return &FileResponseItem{}
}

// FileResponse结构体添加方法
func (f *FileResponse) AddItems(items ...*FileResponseItem) {
	f.Files = append(f.Files, items...)
}

// DeleteFileResponse结构体
type DeleteFileResponse struct {
	Message string `json:"message"`
}

// DeleteFileResponse结构体构造函数
func NewDeleteFileResponse() *DeleteFileResponse {
	return &DeleteFileResponse{}
}

// RetrieveFileContentResponse结构体
type RetrieveFileContentResponse struct {
	Message string `json:"message"`
}

// RetrieveFileContentResponse结构体构造函数
func NewRetrieveFileContentResponse() *RetrieveFileContentResponse {
	return &RetrieveFileContentResponse{}
}
