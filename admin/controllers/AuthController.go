package controllers

import (
	"blog-admin/common"
	"blog-admin/dto"
	"blog-admin/initialization"
	"blog-admin/models"
	"blog-admin/tools"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthController struct{}

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
	// 发放token
	token, err := common.ReleaseToken(admin)
	if err != nil {
		tools.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error: %v", err)
		return
	}
	tools.Success(ctx, gin.H{
		"token": token,
	}, "登录成功")
}

func (receiver AuthController) Info(ctx *gin.Context) {
	admin, _ := ctx.Get("admin")
	tools.Response(ctx, http.StatusOK, 200, gin.H{
		"info": dto.ToAdminDto(admin.(models.Admin)),
	}, "获取成功")
}

func (receiver AuthController) Update(ctx *gin.Context) {
	//admin, _ := ctx.Get("admin")
	DB := initialization.GetDB()
	//info := dto.ToAdminDto(admin.(models.Admin)) // 解析数据
	// 获取参数
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	avatar := ctx.PostForm("avatar")
	autograph := ctx.PostForm("autograph")
	// 更新数据
	var data models.Admin
	DB.First(&data)
	data.UserName = username
	data.Password = password
	data.Email = email
	data.Avatar = avatar
	data.Autograph = autograph
	DB.Save(&data)
	tools.Response(ctx, http.StatusOK, 200, gin.H{
		"info": data,
	}, "更新成功")
}
