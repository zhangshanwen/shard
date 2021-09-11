package model

import (
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	BaseModel
	Username      string `json:"username"     gorm:"size:20"`
	Password      string `json:"-"            gorm:"size:255"`
	Avatar        string `json:"avatar"       gorm:"size:255"`
	RoleId        int64  `json:"role_id"      gorm:"index"`
	Role          Role   `json:"role"         gorm:"foreignkey:RoleId;rerences:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LastLoginTime int64  `json:"last_login_time"`
}

func (u *Admin) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *Admin) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
