package conf

type Conf struct {
	Host          string        `yaml:"host"`
	Rtmp          string        `yaml:"rtmp"`
	Port          string        `yaml:"port"`
	DB            DB            `yaml:"db"`
	Oss           Oss           `yaml:"oss"`
	Authorization Authorization `yaml:"authorization"`
	Level         string        `yaml:"level"`
	ResetPassword string        `yaml:"resetPassword"`
	File          File          `yaml:"file"`
}
