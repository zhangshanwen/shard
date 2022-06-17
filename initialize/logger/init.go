package logger

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Writer io.Writer

func InitGinLogger() {
	gin.DefaultWriter = Writer
	logrus.Info("......GIN日志初始化成功......")
}

func InitLog() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	Writer = io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "log/shard.log",
		MaxSize:    1024, // megabytes
		MaxBackups: 10,
		MaxAge:     7, // days
	})
	logrus.SetOutput(Writer)
	InitGinLogger()
}
