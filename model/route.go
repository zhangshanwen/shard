package model

type Route struct {
	BaseModel
	Method string `json:"method"  gorm:"size:20"` // 方法
	Path   string `json:"path"    gorm:"size:50"` // 路由
}
