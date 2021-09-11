package tools

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type QiNiuImage struct {
	Mac    *qbox.Mac
	domain string
}

func NewQiNiuImage(accessKey, secretKey, domain string) (q *QiNiuImage) {
	return &QiNiuImage{
		Mac:    qbox.NewMac(accessKey, secretKey),
		domain: domain,
	}

}
func (q *QiNiuImage) NewToken(bucket string) string {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	return putPolicy.UploadToken(q.Mac)
}

func (q *QiNiuImage) GetUrl(key string) string {
	return storage.MakePublicURL(q.domain, key)
}
