package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"g1/auth/common"
	"g1/auth/model"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			context.JSON(http.StatusUnauthorized, gin.H{"code": "401", "msg": "权限不足"})
			context.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claim, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"code": "401", "msg": "权限不足"})
			context.Abort()
			return
		}

		userId := claim.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		if user.ID == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{"code": "401", "msg": "权限不足"})
			context.Abort()
			return
		}
		context.Set("user", &user)
		context.Next()
	}
}
