package param

type (
	CheckRoute struct {
		Path   string `json:"path"`
		Method string `json:"method"`
	}
	RouteEdit struct {
		Name string `json:"name"  binding:"required"`
	}
)
