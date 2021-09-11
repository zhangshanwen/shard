package model

type (
	// 频繁修改查询
	Goods struct {
		BaseModel
		Name       string `json:"name"`        // 名称
		Intro      string `json:"intro"`       // 简介
		CategoryId int64  `json:"category_id"` // 分类id
		Cover      string `json:"cover"`       // 封面图
		Inventory  int    `json:"inventory"`   // 库存
		SellPrice  int    `json:"sell_price"`  // 实际售价(单位/分)
		TagId      int64  `json:"tag_id"`      // 标签id
		State      int    `json:"state"`       // 状态 -1-下架 0审核中 1-上架
		Rank       int    `json:"rank"`        // 排序字段越大越靠前
	}
	GoodsDetail struct {
		BaseModel
		GoodsId  int64  `json:"goods_id"`
		Price    int    `json:"price"`    // 售价(单位/分)
		Detail   string `json:"detail"`   // 详情
		Carousel string `json:"carousel"` // 轮播图
	}
)
