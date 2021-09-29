package model

type (
	Carousel struct {
		BaseModel
		Name        string `json:"name"`         // 名称
		ImageId     int64  `json:"image_id"`     // oss id
		RedirectUrl string `json:"redirect_url"` // 跳转url
		Rank        int    `json:"rank"`         // 排序字段越大越靠前
		State       int    `json:"state"`        // 状态 -1-下架 0审核中 1-上架
	}
)
