package param

type (
	UriId struct {
		Id int64 `uri:"id" binding:"required"`
	}
	UriAuthorization struct {
		Authorization string `uri:"authorization"`
	}
	UriStrId struct {
		Id string `uri:"id" binding:"required"`
	}
)
