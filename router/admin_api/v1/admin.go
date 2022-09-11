package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/shard/admin_api/v1/admin"
	"github.com/zhangshanwen/shard/common"
)

func InitAdmin(Router *gin.RouterGroup) {
	r := Router.Group(common.Admins)
	{
		r.POST(common.UriLogin, vt(admin.Login))           // 登录
		r.GET(common.UriEmpty, jwtTx(admin.Get))           // 获取所有管理员
		r.POST(common.UriEmpty, jwtTx(admin.Create))       // 创建管理员
		r.PUT(common.UriId, jwtTx(admin.Edit))             // 修改管理员信息
		r.DELETE(common.UriId, jwtTx(admin.Delete))        // 删除管理员
		r.PUT(common.UriAvatar, jwtTx(admin.UploadAvatar)) // 上传头像

		role := r.Group(common.Roles)
		{
			change := role.Group(common.Change)
			{
				change.PUT(common.UriId, jwtTx(admin.ChangeRole)) // 修改角色
			}
		}

		password := r.Group(common.Password)

		{
			password.PUT(common.Change, jwtTx(admin.ChangePassword)) // 修改密码
			reset := password.Group(common.Reset)
			{
				reset.GET(common.UriId, jwtTx(admin.ResetPassword)) // 重置密码
			}
		}
	}
}
