package controllers

import (
	"blog-admin/common"
	"blog-admin/initialization"
	"blog-admin/models"
	"blog-admin/tools"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

}
