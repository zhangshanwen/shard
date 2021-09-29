package conf

type Oss struct {
	AccessKey  string `yaml:"accessKey"`
	SecretKey  string `yaml:"secretKey"`
	Domain     string `yaml:"domain"`
	AdminBuket string `yaml:"adminBuket"`
}
