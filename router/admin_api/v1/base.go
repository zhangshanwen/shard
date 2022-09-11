package v1

import "github.com/zhangshanwen/shard/middleware"

var (
	v      = middleware.AdminHandel
	vt     = middleware.AdminTxHandel
	jwt    = middleware.AdminJwtHandel
	jwtTx  = middleware.AdminJwtTxHandel
	socket = middleware.AdminSocketHandel
)
