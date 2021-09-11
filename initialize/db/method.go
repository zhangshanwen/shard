package db

import (
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/internal/response"
	"gorm.io/gorm"
)

func FindByPagination(tx *gorm.DB, p *param.Pagination, r *response.Pagination) (err error) {
	r.PageIndex = p.PageIndex
	r.PageSize = p.PageSize
	if err = tx.Count(&r.Total).Error; err != nil {
		return
	}
	tx = tx.Offset(p.Offset()).Limit(p.GetPageSize()).Order(p.OrderBy())
	return
}

func Paginate(p *param.Pagination, r *response.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		r.PageIndex = p.PageIndex
		r.PageSize = p.PageSize
		return db.Offset(p.Offset()).Limit(p.GetPageSize()).Order(p.OrderBy())
	}
}
