package model

// File 文件类型 默认0->可执行文件;1->.py文件;2->python包
type (
	File struct {
		BaseModel
		Hash string `json:"hash"` // 文件hash
		Path string `json:"path"` // 文件路径
	}
	FileRecord struct {
		BaseModel
		Uid      int64  `json:"uid"`
		Name     string `json:"name"`                       // 用户上传文件名用作显示
		FileType uint8  `json:"file_type" gorm:"default:0"` // 文件类型
		FileId   int64  `json:"fileId"`
		File     *File  `gorm:"foreignkey:FileId;rerences:Id;"`
	}
)

func (f *FileRecord) GetCmd() string {
	switch f.FileType {
	case 1:
		return "python3"

	default:
		return "/bin/bash"
	}

}
