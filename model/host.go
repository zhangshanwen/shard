package model

type (
	ConnectType int
	Host        struct {
		BaseModel
		Name        string      `json:"name"`         // 主机名称
		Host        string      `json:"host"`         // 主机地址
		ConnectType ConnectType `json:"connect_type"` // 连接类型 0 http 1 https 2 ssh
		Method      string      `json:"method"`       // 请求方法
		Port        int         `json:"port"`         // 端口
		Username    string      `json:"username"`     // 用户名
		Password    string      `json:"password"`     // 密码
		Status      int         `json:"status"`       // 状态 0 连接失败 1 连接成功
		Comment     string      `json:"comment"`      // 备注
	}
)

const (
	ConnectTypeHttp ConnectType = iota
	ConnectTypeHttps
	ConnectTypeSSh
)
