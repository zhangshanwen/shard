package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Username      string  `json:"username"     gorm:"size:20"`
	Password      string  `json:"-"            gorm:"size:255"`
	Avatar        string  `json:"avatar"       gorm:"size:255"`
	LastLoginTime int64   `json:"last_login_time"`
	Wallet        *Wallet `gorm:"foreignkey:Uid;rerences:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
