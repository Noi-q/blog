package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Avatar    string `json:"avatar"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Autograph string `json:"autograph"`
}

func (Admin) TableName() string {
	return "admin"
}
