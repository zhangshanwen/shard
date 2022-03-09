package db

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"

	"github.com/zhangshanwen/shard/initialize/conf"
)

var R *redis.Client

func InitRedis() {
	logrus.Info("--------init_redis_client_start---------")
	redisConf := conf.C.DB.Redis
	host := redisConf.Host
	port := redisConf.Port
	if host == "" {
		host = "localhost"
	}
	if port <= 0 {
		port = 6379
	}
	R = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", host, port),
		Password: redisConf.Password,
		DB:       redisConf.DataBase,
	})
	if err := R.Ping(context.Background()).Err(); err != nil {
		logrus.Panic(err)
	}
	logrus.Info("--------init_redis_client_end---------")
}
