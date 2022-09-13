package logger

import (
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/zhangshanwen/shard/common"
)

var Writer io.Writer

func InitGinLogger() {
	gin.DefaultWriter = Writer
	logrus.Info("......GIN日志初始化成功......")
}

func InitLog() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: common.TimeFullFormat,
	})
	logrus.SetReportCaller(true)
	l := &lumberjack.Logger{
		Filename:   "log/shard.log",
		MaxSize:    1024, // megabytes
		MaxBackups: 10,
		MaxAge:     7, // days
		LocalTime:  true,
	}
	Writer = io.MultiWriter(os.Stdout, l)

	logrus.SetOutput(Writer)
	InitGinLogger()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)
	go func() {
		for {
			<-c
			l.Rotate()
		}
	}()
}
