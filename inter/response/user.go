package response

import (
	"encoding/json"
	"github.com/jinzhu/copier"
)

type (
	UserInfo struct {
		Authorization string `json:"authorization,omitempty"`
		Avatar        string `json:"avatar"`
		Username      string `json:"username"`
		Balance       int64  `json:"balance"`
	}
	UsersResponse struct {
		List       []User     `json:"list"`
		Pagination Pagination `json:"pagination"`
	}
	User struct {
		Id            int     `json:"id"`
		CreatedTime   int64   `json:"created_time"`
		UpdatedTime   int64   `json:"updated_time"`
		LastLoginTime int64   `json:"last_login_time"`
		Username      string  `json:"username"`
		Wallet        *Wallet `json:"wallet"`
	}
	Wallet struct {
		Balance int64 `json:"balance"`
	}
)

func (p *User) MarshalJSON() ([]byte, error) {

	temp := struct {
		Id            int    `json:"id"`
		CreatedTime   int64  `json:"created_time"`
		UpdatedTime   int64  `json:"updated_time"`
		LastLoginTime int64  `json:"last_login_time"`
		Username      string `json:"username"`
		Balance       int64  `json:"balance"`
	}{}
	_ = copier.Copy(&temp, p)
	if p.Wallet != nil {
		temp.Balance = p.Wallet.Balance
	}
	return json.Marshal(&temp)
}
