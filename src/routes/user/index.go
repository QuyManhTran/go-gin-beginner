package album

import (
	UserController "app/src/controllers/user"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	api := router.Group("/users")
	{
		api.GET("", UserController.GetUsers)
		api.POST("", UserController.CreateUser)
		api.PUT("/:id", UserController.EditUser)
		api.DELETE("/:id", UserController.DeleteUser)
	}
}