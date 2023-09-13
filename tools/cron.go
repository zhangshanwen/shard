package tools

import (
	"github.com/robfig/cron/v3"
	"sync"
)

var (
	c        *cron.Cron
	cronOnce = sync.Once{}
)

func NewCron() *cron.Cron {
	cronOnce.Do(func() {
		c = cron.New()
	})
	return c
}
