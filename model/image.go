package model

type Image struct {
	BaseModel
	Name string `json:"name"`
	Type string `json:"type"` // 类型
}
