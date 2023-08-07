package reply_bot

import (
	"strings"

	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

// Create 创建新的回复机器人
func Create(c *service.AdminTxContext) (r service.Res) {
	p := param.SaveReplyBot{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		tx = c.Tx
		m  = model.ReplyBot{
			Name: p.Name,
			Uid:  c.Admin.Id,
		}
	)
	var count int64
	if r.Err = tx.Model(&m).Where(&m).Count(&count).Error; r.Err != nil {
		return
	}
	if count > 0 {
		r.NameIsExisted()
		return
	}
	if r.Err = copier.Copy(&m, &p); r.Err != nil {
		r.CopierError()
		return
	}
	m.Friends = strings.Join(p.Friends, ",")
	for _, i := range p.RuleIds {
		m.Rules = append(m.Rules, model.Rule{
			BaseModel: model.BaseModel{Id: i},
		})
	}
	if r.Err = tx.Create(&m).Error; r.Err != nil {
		r.DBError()
		return
	}
	return
}
