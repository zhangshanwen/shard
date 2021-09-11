package param

type (
	AdminRegister struct {
		Username string `json:"username"  binding:"required"`
		Password string `json:"password"   binding:"required"`
	}
	AdminLogin struct {
		Username string `json:"username"  binding:"required"`
		Password string `json:"password"   binding:"required"`
	}
	AdminChangePassword struct {
		Password    string `json:"password"       binding:"required"`
		NewPassword string `json:"new_password"   binding:"required"`
	}
	AdminRecords struct {
		Username string `form:"username"`
		Pagination
	}
	AdminEdit struct {
		Username string `json:"username"  binding:"required"`
	}
	AdminUploadAvatar struct {
		Avatar string `json:"avatar"  binding:"required"`
	}
	AdminChangeRole struct {
		RoleId int64 `json:"role_id"`
	}
)
