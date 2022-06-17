package response

type (
	FileResponse struct {
		List       []File     `json:"list"`
		Pagination Pagination `json:"pagination"`
	}
	File struct {
		Id          int64  `json:"id"`
		CreatedTime int64  `json:"created_time"`
		UpdatedTime int64  `json:"updated_time"`
		Name        string `json:"name"`
		FileType    uint8  `json:"file_type"`
	}
	FileDetail struct {
		Id       int64  `json:"id"`
		Name     string `json:"name"`
		Code     string `json:"code"`
		FileType uint8  `json:"file_type"`
	}
)
