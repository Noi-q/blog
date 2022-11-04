package router

import (
	"blog-admin/controllers"
	"blog-admin/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine) {
	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"result": "200",
		})
	})

	auth := engine.Group("/api/auth/")
	{
		auth.POST("/")
		auth.POST("/login", controllers.AuthController{}.Login)                                     // 登录
		auth.POST("/info", middleware.AuthMiddleware(), controllers.AuthController{}.Info)          // 信息
		auth.POST("/update_info", middleware.AuthMiddleware(), controllers.AuthController{}.Update) // 更新
	}

	// 文件上传
	engine.POST("/api/upload", middleware.AuthMiddleware(), controllers.UploadController{}.Upload)
}
