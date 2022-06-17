package response

type (
	HostResponse struct {
		List       []Host     `json:"list"`
		Pagination Pagination `json:"pagination"`
	}
	Host struct {
		Id          int64  `json:"id"`
		CreatedTime int64  `json:"created_time"`
		UpdatedTime int64  `json:"updated_time"`
		Name        string `json:"name"`         // 主机名称
		Host        string `json:"host"`         // 主机地址
		Username    string `json:"username"`     // 账号
		Password    string `json:"password"`     // 密码
		ConnectType int    `json:"connect_type"` // 连接类型 0 http/https 1 ssh
		Port        int    `json:"port"`         // 端口
		Status      int    `json:"status"`       // 状态 0未尝试连接 -1 连接失败 1 连接成功
		Comment     string `json:"comment"`      // 备注
	}
	Room struct {
		Id string `json:"id"`
	}
)
