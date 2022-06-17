package response

type (
	RouteResponse struct {
		List []Route `json:"list"`
	}
	Route struct {
		Id          int64  `json:"id"`
		CreatedTime int64  `json:"created_time"`
		UpdatedTime int64  `json:"updated_time"`
		Method      string `json:"method"`
		Path        string `json:"path"`
	}
)
