package auth

import (
	AuthController "app/src/controllers/auth"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine) {
	api := router.Group("/auth")
	{
		api.POST("/login", AuthController.Login)
	}
}