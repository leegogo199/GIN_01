package middleware

import (
	"GINVUE/Model"
	"GINVUE/common"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized,
				gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return

		}
		//
		tokenString = tokenString[7:]
		//
		token, claims, err := common.ParseToken(tokenString)
		//
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//验证通过后获取claim中的userID
		userId := claims.UserId
		DB := common.GetDB()
		var user Model.User
		DB.First(&user, userId)
		//用户不存在
		if userId == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//用户存在，将user的信息写入上下午。
		ctx.Set("user", user)
		ctx.Next()
	}
}
