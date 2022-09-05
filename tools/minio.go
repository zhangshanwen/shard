package tools

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/qiniu/go-sdk/v7/storage"
)

type MinioOss struct {
	Mac    *minio.Client
	domain string
}

func NewMinioImage(accessKey, secretKey, domain string) (q *MinioOss) {
	c, _ := minio.New(domain, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	return &MinioOss{
		Mac:    c,
		domain: domain,
	}

}
func (q *MinioOss) NewToken(bucket string) (t string) {
	return
}

func (q *MinioOss) GetUrl(key string) string {
	return storage.MakePublicURL(q.domain, key)
}
