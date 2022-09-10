package tools

import (
	"context"
	"sync"

	"github.com/zhangshanwen/shard/initialize/conf"
)

type Oss interface {
	NewToken(bucket string) string
	GetUrl(ctx context.Context, key string) string
	UploadFile(ctx context.Context, filename string, file []byte) (string, error)
}

var (
	o    Oss
	once sync.Once
)

func NewOss() (Oss, error) {
	var err error
	if o == nil {
		once.Do(func() {
			o, err = NewMinioImage(conf.C.Oss.AccessKey, conf.C.Oss.SecretKey, conf.C.Oss.Domain)
		})
	}
	return o, err

}
