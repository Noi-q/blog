package router

import (
	"blog-admin/controllers"
	"blog-admin/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine) {
	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"code": 200,
			"msg":  "接口测试",
		})
	})
	//
	auth := engine.Group("/api/auth")
	{
		auth.POST("/login", controllers.AuthController{}.Login)                                     // 登录
		auth.POST("/info", middleware.AuthMiddleware(), controllers.AuthController{}.Info)          // 信息
		auth.POST("/update_info", middleware.AuthMiddleware(), controllers.AuthController{}.Update) // 更新
	}
	// 文件上传
	upload := engine.Group("/api/upload")
	{
		upload.POST("/avatar", middleware.AuthMiddleware(), controllers.UploadController{}.Img) // 头像上传
	}

	// 栏目
	category := engine.Group("/api/category")
	{
		category.POST("/insert", middleware.AuthMiddleware(), controllers.CategoryController{}.Insert) // 添加栏目
		category.POST("/update", middleware.AuthMiddleware(), controllers.CategoryController{}.Update) // 修改栏目
		category.POST("/delete", middleware.AuthMiddleware(), controllers.CategoryController{}.Delete) // 删除栏目
		category.GET("/query", middleware.AuthMiddleware(), controllers.CategoryController{}.Query)    // 查询栏目
	}
	// 文章
	article := engine.Group("/api/article")
	{
		article.POST("/insert", middleware.AuthMiddleware(), controllers.ArticleController{}.Insert) // 添加文章
		article.POST("/update", middleware.AuthMiddleware(), controllers.ArticleController{}.Update) // 修改文章
		article.POST("/delete", middleware.AuthMiddleware(), controllers.ArticleController{}.Delete) // 删除文章
		article.GET("/query", middleware.AuthMiddleware(), controllers.ArticleController{}.Query)    // 查询文章
	}
}
