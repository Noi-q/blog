package dto

import "blog-admin/models"

type AdminDto struct {
	UserName  string
	Email     string
	Avatar    string
	Autograph string
}

func ToAdminDto(admin models.Admin) AdminDto {
	return AdminDto{
		UserName:  admin.UserName,
		Email:     admin.Email,
		Avatar:    admin.Avatar,
		Autograph: admin.Autograph,
	}
}
