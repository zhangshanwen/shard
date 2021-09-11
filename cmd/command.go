package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zhangshanwen/shard/initialize/conf"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/logger"
	"github.com/zhangshanwen/shard/model"
)

var usageTemplate = `backend commands 
go run cmd/command.go  create_admin -username '' -password ''

`
var (
	username = flag.String("username", "", "please input your username")
	password = flag.String("password", "", "please input your password ")
)

func Usage() {
	_, _ = fmt.Fprintf(flag.CommandLine.Output(), usageTemplate)
	flag.PrintDefaults()
}
func main() {

	args := os.Args
	if len(args) < 2 {
		UsageAndExit()
		return
	}
	cmd := args[1]
	os.Args = os.Args[1:]
	switch cmd {
	case "help":
		UsageAndExit()
	case "create_admin":
		CreateAdmin()
	default:
		PrintErrorAndExit("Unknown subcommand")
	}
}
func UsageAndExit() {
	Usage()
	os.Exit(2)
	return
}

func PrintErrorAndExit(message string) {
	fmt.Println(message)
	os.Exit(0)
}

func CreateAdmin() {
	flag.Parse()

	if *username == "" {
		PrintErrorAndExit("please input your username")
		return
	}
	if *password == "" {
		PrintErrorAndExit("please input your password")
		return
	}
	// 创建数据链接
	logger.InitLog() // 初始化日志
	conf.InitConf()  // 初始化配置文件
	db.InitMysql()
	admin := model.Admin{Username: *username}
	var count int64
	if err := db.G.Model(&model.Admin{}).Where(&admin).Count(&count).Error; err != nil {
		PrintErrorAndExit(err.Error())
		return
	}
	if count > 0 {
		PrintErrorAndExit("username is existed")
		return
	}
	if err := admin.SetPassword(*password); err != nil {
		PrintErrorAndExit(err.Error())
		return
	}
	if err := db.G.Save(&admin).Error; err != nil {
		PrintErrorAndExit(err.Error())
		return
	}
	fmt.Printf("create username=%v password=%v success \n ", *username, *password)
}
