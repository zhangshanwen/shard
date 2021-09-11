package response

type (
	UserInfo struct {
		Authorization string `json:"authorization,omitempty"`
		Avatar        string `json:"avatar"`
		Username      string `json:"username"`
		Balance       int64  `json:"balance"`
	}
	UsersResponse struct {
		List       []User     `json:"list"`
		Pagination Pagination `json:"pagination"`
	}
	User struct {
		Id            int    `json:"id"`
		CreatedTime   int64  `json:"created_time"`
		UpdatedTime   int64  `json:"updated_time"`
		LastLoginTime int64  `json:"last_login_time"`
		Username      string `json:"username"`
	}
)
