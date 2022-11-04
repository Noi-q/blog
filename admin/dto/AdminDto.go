package dto

import (
	"blog-admin/models"
)

type AdminDto struct {
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Autograph string `json:"autograph"`
}

func ToAdminDto(admin models.Admin) AdminDto {
	return AdminDto{
		UserName:  admin.UserName,
		Email:     admin.Email,
		Avatar:    admin.Avatar,
		Autograph: admin.Autograph,
	}
}
