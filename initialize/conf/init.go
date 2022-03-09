package conf

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	C *Conf
)

const (
	ConfigFile = "./env/env.yaml"
	ConfigEnv  = "TheOneConfig"
)

func InitConf(path ...string) {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			config = ConfigFile
			if configEnv := os.Getenv(ConfigEnv); configEnv == "" {
				config = ConfigFile
			} else {
				config = configEnv
				logrus.Info("您正在使用GVA_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			logrus.Info("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		logrus.Info("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		logrus.Fatal("Fatal error config file: %s \n", err)
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&C); err != nil {
			logrus.Fatal(err)
		}
	})
	if err := v.Unmarshal(&C); err != nil {
		logrus.Fatal(err)
	}
	return
}
