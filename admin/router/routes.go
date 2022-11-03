package router

import (
	"blog-admin/controllers"
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
		auth.POST("/login", controllers.AuthController{}.Login)
		auth.POST("/info", controllers.AuthController{}.Info)
		auth.POST("/update_password")
	}
}
