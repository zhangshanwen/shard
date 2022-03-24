package param

type (
	FileUploadParams struct {
		File     string `json:"file"        binding:"required"`
		FileName string `json:"file_name"   binding:"required"`
		FileType uint8  `json:"file_type"   `
	}
	FileRecords struct {
		Pagination
	}
	FileRunParams struct {
		Id int64 `json:"id" binding:"required"`
	}
)
