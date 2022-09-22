package param

type (
	LogRecords struct {
		Pagination
		Types string `form:"types"`
	}
	LogDel struct {
		Ids string `form:"ids"`
	}
)
