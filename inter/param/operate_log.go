package param

type (
	LogRecords struct {
		Pagination
		Types string `form:"types"`
	}
)
