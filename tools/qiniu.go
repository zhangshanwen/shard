package tools

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type QiNiuiOss struct {
	Mac    *qbox.Mac
	domain string
}

func NewQiNiuOss(accessKey, secretKey, domain string) (q *QiNiuiOss) {
	return &QiNiuiOss{
		Mac:    qbox.NewMac(accessKey, secretKey),
		domain: domain,
	}

}
func (q *QiNiuiOss) NewToken(bucket string) string {
	putPolicy := storage.PutPolicy{
		Scope:   bucket,
		Expires: 7200,
	}
	return putPolicy.UploadToken(q.Mac)
}

func (q *QiNiuiOss) GetUrl(key string) string {
	return storage.MakePublicURL(q.domain, key)
}
