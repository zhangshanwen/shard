package model

type (
	GoodsCategory struct {
		BaseModel
		Name     string `json:"name"`      // 名称
		ParentId int64  `json:"parent_id"` // 父分类id
	}
)
