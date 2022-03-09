package logger

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Writer io.Writer

func InitGinLogger() {
	gin.DefaultWriter = Writer
	log.Info("......GIN日志初始化成功......")
}

func InitLog() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})
	Writer = io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "log/shard.log",
		MaxSize:    1024, // megabytes
		MaxBackups: 10,
		MaxAge:     7, // days
	})
	log.SetOutput(Writer)
	InitGinLogger()
}
