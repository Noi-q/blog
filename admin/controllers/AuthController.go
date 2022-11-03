package controllers

import (
	"blog-admin/dto"
	"blog-admin/initialization"
	"blog-admin/models"
	"blog-admin/tools"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
}

func (receiver AuthController) Login(ctx *gin.Context) {
	// 获取参数
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	// 查询数据
	var admin models.Admin
	db := initialization.GetDB()
	db.Where("user_name = ?", username).Where("password = ?", password).First(&admin)
	// 判断用户是否存在
	if admin.ID == 0 {
		tools.Fail(ctx, 400, nil, "用户或密码错误")
		return
	}
	tools.Success(ctx, gin.H{
		"data": dto.ToAdminDto(admin),
	}, "登录成功")
}

func (receiver AuthController) Info(ctx *gin.Context) {

}
