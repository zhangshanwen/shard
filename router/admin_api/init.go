package admin_api

import (
	"github.com/zhangshanwen/shard/common"
	"github.com/zhangshanwen/shard/initialize/app"
	v1 "github.com/zhangshanwen/shard/router/admin_api/v1"
)

func RegisterBackendV1Router() {
	api := app.R.Group(common.BackendPrefix)
	group := api.Group(common.V1)

	v1.InitVersion(group)
	v1.InitAdmin(group)
	v1.InitPermission(group)
	v1.InitUser(group)
	v1.InitRoute(group)
	v1.InitRole(group)
	v1.InitOss(group)
	v1.InitFile(group)
	v1.InitTask(group)
	v1.InitHost(group)
	v1.InitOperateLog(group)
	v1.InitLive(group)
	v1.InitWechat(group)
	v1.InitMeeting(group)
}
