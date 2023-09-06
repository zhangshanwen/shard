package reply_bot

import (
	"strings"

	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

func Edit(c *service.AdminTxContext) (r service.Res) {
	pId := param.UriId{}
	if r.Err = c.BindUri(&pId); r.Err != nil {
		r.ParamsError()
		return
	}
	p := param.SaveReplyBot{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		m  = model.ReplyBot{}
		tx = c.Tx
	)
	defer func() {
		if r.Err == nil {
			r.Data = m
		}
	}()
	if r.Err = tx.First(&m, pId.Id).Error; r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = copier.Copy(&m, &p); r.Err != nil {
		r.CopierError()
		return
	}
	var rules []model.Rule
	m.Friends = strings.Join(p.Friends, ",")
	m.Groups = strings.Join(p.Groups, ",")
	for _, routeId := range p.RuleIds {
		rules = append(rules, model.Rule{
			BaseModel: model.BaseModel{Id: routeId},
		})
	}
	if r.Err = tx.Model(&m).Association("Rules").Replace(&rules); r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = tx.Save(&m).Error; r.Err != nil {
		r.DBError()
		return
	}
	return
}
