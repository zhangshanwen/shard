package rules

import (
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/tools/wechat"
)

// Functions 返回默认返回规则
func Functions(c *service.AdminContext) (r service.Res) {
	var (
		resp response.DefaultRulesResponse
	)
	for functionName, v := range wechat.DefaultTemplateReply {
		resp.List = append(resp.List, response.DefaultRule{
			FunctionName: functionName, Description: v.Description,
		})
	}
	r.Data = resp
	return
}
