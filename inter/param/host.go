package param

type (
	HostSave struct {
		Name        string `json:"name"         binding:"required"`        // 主机名称
		Host        string `json:"host"`                                   // 主机地址
		Username    string `json:"username"`                               // 账号
		Password    string `json:"password"`                               // 密码
		ConnectType int    `json:"connect_type" binding:"gte=0,lte=2"`     // 连接类型 0 http 1 https 2 ssh
		Port        int    `json:"port"         binding:"gte=0,lte=65535"` // 端口
		Comment     string `json:"comment"       `                         // 备注
	}
	Host struct {
		Pagination
	}
	Room struct {
		Id int64 `json:"id" binding:"required"`
	}
	Socket struct {
		Id string `uri:"id" binding:"required"`
	}
	SocketInt struct {
		Id int64 `uri:"id" binding:"required"`
	}
)
