package v1

import "github.com/zhangshanwen/shard/middleware"

var (
	v      = middleware.AdminHandel
	jwt    = middleware.AdminJwtHandel
	socket = middleware.AdminSocketHandel
)
