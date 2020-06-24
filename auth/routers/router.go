package routers

import (
	"github.com/gin-gonic/gin"

	"g1/auth/controller"
	"g1/auth/middleware"
)

func Routes(r *gin.Engine) {
	user := r.Group("/api/auth")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
		user.GET("/info", middleware.Auth(), controller.Info)
	}
}