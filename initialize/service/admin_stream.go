package service

import "github.com/zhangshanwen/shard/live/av"

type (
	AdminStreamContext struct {
		AdminTxContext
		Handler av.Handler
	}
)
