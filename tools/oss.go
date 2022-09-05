package tools

import (
	"sync"

	"github.com/zhangshanwen/shard/initialize/conf"
)

type Oss interface {
	NewToken(bucket string) string
	GetUrl(key string) string
}

var (
	o    Oss
	once sync.Once
)

func NewOss() Oss {
	if o == nil {
		once.Do(func() {
			o = NewQiNiuOss(conf.C.Oss.AccessKey, conf.C.Oss.SecretKey, conf.C.Oss.Domain)
		})
	}

	return o
}
