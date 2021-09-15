package router

import (
	"fmt"

	"github.com/zhangshanwen/shard/initialize/app"
	"github.com/zhangshanwen/shard/initialize/conf"
	"github.com/zhangshanwen/shard/router/admin_api"
	"github.com/zhangshanwen/shard/router/api"
)

func InitRouter() {
	api.RegisterApiV1Router()
	admin_api.RegisterBackendV1Router()
	run()
}

func run() {
	app.InitRoute()
	_ = app.R.Run(fmt.Sprintf("%s:%s", conf.C.Host, conf.C.Port))

}
