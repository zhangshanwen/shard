package conf

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	C       *Conf
	Project string
)

const (
	ConfigFile = "./env/env.yaml"
)

func InitConf() {
	v := viper.New()
	v.SetConfigFile(ConfigFile)
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
