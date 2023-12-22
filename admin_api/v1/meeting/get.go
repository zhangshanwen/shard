package meeting

import (
	"fmt"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
)

func Get(c *service.AdminTxContext) (r service.Res) {
	p := param.MeetingRecords{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.DBError()
		return
	}
	var (
		m    model.Meeting
		ms   []model.Meeting
		tx   = c.Tx
		resp = response.MeetingResponse{}
	)
	defer func() {
		if r.Err == nil {
			r.Data = resp
		}
	}()
	g := tx.Model(&m)
	if p.Name != "" {
		m.Name = fmt.Sprintf("%%%s%%", p.Name)
		g = g.Where(&m)
	}

	if r.Err = db.FindByPagination(g, &p.Pagination, &resp.Pagination); r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = g.Preload("Owner").Find(&ms).Error; r.Err != nil {
		r.DBError()
		return
	}
	for _, item := range ms {
		resp.List = append(resp.List, response.Meeting{
			Id:          item.Id,
			Name:        item.Name,
			Status:      int(item.Status),
			CreatedTime: item.CreatedTime,
			UpdatedTime: item.UpdatedTime,
			StartTime:   item.StartTime,
			EndTime:     item.EndTime,
			IsOwner:     item.Owner.Id == c.Admin.Id,
			Owner:       item.Owner.Username,
		})
	}
	return
}
