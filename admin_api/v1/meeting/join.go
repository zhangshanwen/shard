package meeting

import (
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/rtmp"
)

func Join(c *service.AdminContext) (r service.Res) {
	var (
		member *rtmp.Member
		p      param.UriId
		m      model.Meeting
	)
	if r.Err = c.BindUri(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	if r.Err = db.G.First(&m, p.Id).Error; r.Err != nil {
		r.RoomJoinFailed()
		return
	}
	if member, r.Err = rtmp.S.AddMember(p.Id, c.Done, c.Writer); r.Err != nil {
		r.RoomJoinFailed()
		return
	}
	defer member.Quit()
	member.Wait()
	return
}
