package model

// File 文件类型 默认0->可执行文件;1->.py文件;2->python包
type (
	File struct {
		BaseModel
		Name string `json:"name"`
		Hash string `json:"hash"`                  // 文件hash
		Path string `json:"path"`                  // 文件路径
		Type uint8  `json:"type" gorm:"default:0"` // 文件类型
	}
	FileRecord struct {
		BaseModel
		Uid    int64 `json:"uid"`
		FileId int64 `json:"fileId"`
	}
)
