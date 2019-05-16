package models

import (
	"github.com/jinzhu/gorm"
	"go-clean-architure/libs/common"
)

type User struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
}

// Serialize serializes user data
func (u *User) Serialize() common.JSON {
	return common.JSON{
		"id":           u.ID,
		"email":     u.Email,
		"password": u.Password,
	}
}

func (u *User) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	u.Email = m["email"].(string)
	u.Password = m["password"].(string)
}