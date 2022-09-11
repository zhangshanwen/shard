package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangshanwen/shard/admin_api/v1/user"
	"github.com/zhangshanwen/shard/common"
)

func InitUser(Router *gin.RouterGroup) {
	r := Router.Group(common.Users)
	{
		r.GET(common.UriEmpty, jwtTx(user.Get))     // 获取所有用户
		r.POST(common.UriEmpty, jwtTx(user.Create)) // 创建用户
		r.PUT(common.UriId, jwtTx(user.Edit))       // 修改用户信息
		r.DELETE(common.UriId, jwtTx(user.Delete))  // 删除用户

		password := r.Group(common.Password)
		{
			reset := password.Group(common.Reset)
			{
				reset.GET(common.UriId, jwtTx(user.ResetPassword)) // 重置密码
			}
		}
		balance := r.Group(common.Balance)
		{
			adjust := balance.Group(common.Adjust)
			{
				adjust.PATCH(common.UriId, jwtTx(user.BalanceAdjust)) // 调整余额
			}
		}
	}
}
