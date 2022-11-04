package middleware

import (
	"blog-admin/common"
	"blog-admin/initialization"
	"blog-admin/models"
	"blog-admin/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取authorization header
		tokenString := context.GetHeader("Authorization")
		// validate token
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			tools.Response(context, http.StatusUnauthorized, 401, nil, "权限不足")
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			tools.Response(context, http.StatusUnauthorized, 401, nil, "权限不足")
			context.Abort()
			return
		}
		// 验证通关后获取claim 中的userId
		userId := claims.UserId
		DB := initialization.GetDB()
		var admin models.Admin
		DB.First(&admin, userId)
		// 用户不存在
		if admin.ID == 0 {
			tools.Response(context, http.StatusUnauthorized, 401, nil, "权限不足")
			context.Abort()
			return
		}
		// 用户存在，将admin的信息写入上下文
		context.Set("admin", admin)
		context.Next()
	}
}
