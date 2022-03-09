package db

func InitDb() {
	InitMysql() // 初始化mysql
	InitRedis() // 初始化redis
}
