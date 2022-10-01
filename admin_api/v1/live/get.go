package live

import (
	"fmt"

	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools"
)

func Get(c *service.AdminTxContext) (r service.Res) {
	p := param.LiveRoom{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ResCode = code.ParamsError
		return
	}
	var (
		tx   = c.Tx
		resp response.LiveRoomResponse
		m    model.LiveRoom
		ms   []model.LiveRoom
	)

	defer func() {
		if r.Err == nil {
			r.Data = resp
		}
	}()
	g := tx.Model(&m)
	if r.Err = db.FindByPagination(g, &p.Pagination, &resp.Pagination); r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = g.Preload("Owner").Find(&ms).Error; r.Err != nil {
		r.DBError()
		return
	}
	for _, i := range ms {
		resp.List = append(resp.List, response.LiveRoom{
			Id:          i.Id,
			Name:        i.Name,
			Owner:       i.Owner.Username,
			Status:      int(i.Status),
			CreatedTime: i.CreatedTime,
			UpdatedTime: i.UpdatedTime,
			StartTime:   i.StartTime,
			EndTime:     i.EndTime,
			Hash:        tools.Hash(fmt.Sprintf("%v_%v", i.Owner.Id, i.Name)),
		})
	}
	c.SaveLogSelect(tx, module, fmt.Sprintf("cat list"))
	return
}
