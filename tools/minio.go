package tools

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"github.com/zhangshanwen/shard/initialize/conf"
	"net/url"
	"time"
)

type MinioOss struct {
	Mac    *minio.Client
	domain string
}

func NewMinioImage(accessKey, secretKey, domain string) (q *MinioOss, err error) {
	q = &MinioOss{
		domain: domain,
		Mac:    &minio.Client{},
	}
	q.Mac, err = minio.New(domain, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	return

}
func (q *MinioOss) NewToken(bucket string) (t string) {
	return
}

func (q *MinioOss) GetUrl(ctx context.Context, key string) (s string) {
	var (
		u   *url.URL
		err error
	)
	if u, err = q.Mac.PresignedGetObject(context.Background(), conf.C.Oss.AdminBuket, key, time.Minute*30, url.Values{}); err != nil {
		return
	}
	return u.String()
}

func (q *MinioOss) UploadFile(ctx context.Context, filename string, file []byte) (key string, err error) {
	var info minio.UploadInfo
	if info, err = q.Mac.PutObject(ctx, conf.C.Oss.AdminBuket, filename, bytes.NewBuffer(file), int64(len(file)), minio.PutObjectOptions{}); err != nil {
		return
	}
	logrus.Info(info)
	logrus.Info("----------")
	return info.Key, nil
}
