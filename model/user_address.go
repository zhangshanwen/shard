package model

type UserAddress struct {
	BaseModel
	Uid         int64  `json:"uid"`
	Province    string `json:"province"`     // 省/直辖市
	City        string `json:"city"`         // 市
	Region      string `json:"region"`       // 区
	NickName    string `json:"nick_name"`    // 别名
	Name        string `json:"name"`         // 收货人名字
	PhoneNumber string `json:"phone_number"` // 收货人地址
	Address     string `json:"address"`      // 详细地址
	PostalCode  string `json:"postal_code"`  // 邮政编码
	Flag        uint8  `json:"flag"`         // 1默认 0 普通
}
