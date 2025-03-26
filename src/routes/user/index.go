package user

import (
	UserController "app/src/controllers/user"
	"app/src/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	api := router.Group("/users")
	api.Use(middlewares.JwtMiddleware())
	{
		api.GET("", UserController.GetUsers)
		api.POST("", UserController.CreateUser)
		api.PUT("/:id", UserController.EditUser)
		api.DELETE("/:id", UserController.DeleteUser)
	}
}