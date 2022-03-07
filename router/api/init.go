package api

import (
	"github.com/zhangshanwen/shard/initialize/app"
	"github.com/zhangshanwen/shard/router/api/v1"
)

func RegisterApiV1Router() {
	api := app.R.Group("api")
	group := api.Group("v1")
	v1.InitVersion(group)
	v1.InitUser(group)
	v1.InitUpload(group)
}
