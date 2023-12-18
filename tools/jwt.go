package tools

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"

	"github.com/zhangshanwen/shard/initialize/conf"
)

var (
	defaultExpiresTimes = 12 * time.Hour
	defaultTokenType    = conf.Project
	privateKey          *rsa.PrivateKey
	publicKey           *rsa.PublicKey
	method              = jwt.SigningMethodRS256 //默认256
)

type Payload struct {
	Uid       int64
	TokenType string
}

type Claims struct {
	*jwt.StandardClaims
	Payload
}

func Load() {
	var err error
	var privateBytes, publicBytes []byte
	privateBytes = []byte(privateKeyStr)
	publicBytes = []byte(publicKeyStr)
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		logrus.Fatalf("[initKeys]: %s\n", err)
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		logrus.Fatalf("[initKeys]: %s\n", err)
	}
}

func GetExpireTime() (ExpireHour time.Duration) {
	if conf.C.Authorization.ExpireHour == 0 {
		return defaultExpiresTimes
	} else {
		return time.Duration(conf.C.Authorization.ExpireHour) * time.Hour
	}
}
func GetExpiresAt() (expiresAt int64) {
	return time.Now().Add(GetExpireTime()).Unix()
}

func CreateToken(uid int64) (token string, err error) {
	t := jwt.NewWithClaims(method, Claims{
		&jwt.StandardClaims{
			ExpiresAt: GetExpiresAt(),
		},
		Payload{uid, defaultTokenType},
	})
	return t.SignedString(privateKey)
}
func VerifyToken(tokenString string) (claims *Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return publicKey, nil
	})
	if err != nil {
		return
	}
	claims = token.Claims.(*Claims)
	return
}
