package param

type Pagination struct {
	PageIndex int    `form:"page_index"`
	PageSize  int    `from:"page_size"`
	Desc      bool   `form:"desc"`
	Order     string `form:"order"`
}

func (p *Pagination) GetPageIndex() int {
	if p.PageIndex == 0 {
		return 1
	}
	return p.PageIndex
}
func (p *Pagination) GetPageSize() int {
	if p.PageSize == 0 {
		return 20
	}
	return p.PageSize
}
func (p *Pagination) Offset() int {
	return (p.GetPageIndex() - 1) * p.GetPageSize()
}

func (p *Pagination) OrderBy() string {
	orderBy := p.Order
	if p.Order == "" {
		orderBy = "id"
	}
	if p.Desc {
		orderBy += " desc"
	}
	return orderBy
}
