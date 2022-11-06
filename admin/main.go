package main

import (
	"blog-admin/common"
	"blog-admin/initialization"
	"blog-admin/models"
	"blog-admin/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	initialization.InitConfig()            // 初始化配置文件
	initialization.InitDatabase()          // 初始化数据库
	CreateRoot()                           // 创建管理员
	r := gin.New()                         // 初始化Gin
	r.Use(common.Logger(), gin.Recovery()) // 初始化自定义日志中间件
	r.Static("/static", "./static")        // 初始化文件资源目录
	router.Routes(r)                       // 初始化路由
	// 配置端口
	if port := viper.GetString("server.port"); port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func CreateRoot() {
	db := initialization.GetDB()
	var admin models.Admin
	db.Where("id = ?", 1).First(&admin)
	if admin.ID != 1 && admin.UserName != viper.GetString("admin.username") {
		db.Create(&models.Admin{
			UserName: viper.GetString("admin.username"),
			Password: viper.GetString("admin.password"),
		})
	}
}
