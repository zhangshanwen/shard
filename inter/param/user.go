package param

type (
	Register struct {
		Username string `json:"username"  binding:"required"`
		Password string `json:"password"   binding:"required"`
	}
	Login struct {
		Username string `json:"username"  binding:"required"`
		Password string `json:"password"   binding:"required"`
	}
	ChangePassword struct {
		Password    string `json:"password"       binding:"required"`
		NewPassword string `json:"new_password"   binding:"required"`
	}
	UserRecords struct {
		Username string `form:"username"`
		Pagination
	}
	UserEdit struct {
		Username string `json:"username"  binding:"required"`
	}
	UserUploadAvatar struct {
		Avatar string `json:"avatar"  binding:"required"`
	}
	AdjustBalance struct {
		Amount int64 `json:"amount"       binding:"required"`
	}
)
