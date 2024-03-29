package initialize

import (
	"github.com/zhangshanwen/shard/initialize/conf"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/logger"
	"github.com/zhangshanwen/shard/initialize/node"
	"github.com/zhangshanwen/shard/initialize/task"
	"github.com/zhangshanwen/shard/tools"
)

func Initialize() {
	logger.InitLog() // 初始化日志
	conf.InitConf()  // 初始化配置文件
	db.InitDb()      // 初始化数据库
	task.InitTask()  // 初始化任务
	node.InitNode()  // 初始化node
	tools.Load()     // 初始化jwt
}
