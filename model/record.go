package model

type (
	Record struct {
		BaseModel
		Uid     int64  `json:"uid"`
		Number  string `json:"number"`   // 订单号
		Amount  int    `json:"amount"`   // 支付总价
		State   int    `json:"state"`    // 支付状态  -2.手动关闭 -3.超时关闭 -4.商家关闭 -1支付失败 0未支付 1 支付成功/待发货 2 配货中 3 配货完成 4.出库成功 5.运输中 6 交易成功
		Type    uint   `json:"type"`     // 支付类型 0 余额 1 支付宝 2 微信 3 ....
		PayTime int    `json:"pay_time"` // 付款时间
	}
	RecordAddress struct {
		RecordId   int64  `json:"record_id"`
		Province   string `json:"province"`    // 省/直辖市
		City       string `json:"city"`        // 市
		Region     string `json:"region"`      // 区
		Address    string `json:"address"`     // 详细地址
		PostalCode string `json:"postal_code"` // 邮政编码
		State      int    `json:"state"`       // 状态 -1废弃 0使用中
	}
)
