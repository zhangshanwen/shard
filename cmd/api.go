package main

import (
	"github.com/zhangshanwen/shard/initialize"
	"github.com/zhangshanwen/shard/router"
)

func main() {
	//rpc.SendServer()
	initialize.Initialize() // 注册服务
	router.InitRouter()     // 注册路由
}
