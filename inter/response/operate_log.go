package response

import "encoding/json"

type (
	LogResponse struct {
		List       []Log      `json:"list"`
		Pagination Pagination `json:"pagination"`
	}
	Log struct {
		Id          int64  `json:"id"`
		CreatedTime int64  `json:"created_time"`
		Log         string `json:"log"`
		Type        uint8  `json:"type"`
		Admin       *Admin
	}
)

func (p *Log) MarshalJSON() ([]byte, error) {
	var (
		username string
		roleName string
	)
	if p.Admin != nil {
		username = p.Admin.Username
		roleName = p.Admin.Role.Name
	}
	temp := struct {
		Id          int64  `json:"id"`
		CreatedTime int64  `json:"created_time"`
		Log         string `json:"log"`
		RoleName    string `json:"role_name"`
		Type        uint8  `json:"type"`
		Username    string `json:"username"`
	}{
		Id:          p.Id,
		CreatedTime: p.CreatedTime,
		Log:         p.Log,
		RoleName:    roleName,
		Type:        p.Type,
		Username:    username,
	}
	return json.Marshal(&temp)
}
