package rules

import (
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/tools/wechat"
)

// DefaultRules 返回默认返回规则
func DefaultRules(c *service.AdminContext) (r service.Res) {
	var (
		resp response.DefaultRulesResponse
	)
	for name, v := range wechat.DefaultTemplateReply {
		resp.List = append(resp.List, response.DefaultRule{
			Name: name, Desc: v.Desc,
		})
	}
	r.Data = resp
	return
}
