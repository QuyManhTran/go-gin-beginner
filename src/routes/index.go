package route

import (
	LoggerMiddleware "app/src/middlewares"
	AuthRouter "app/src/routes/auth"
	UserRouter "app/src/routes/user"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine, port string) {
	router.Use(LoggerMiddleware.Logger())
	UserRouter.UserRouter(router)
	AuthRouter.AuthRouter(router)
	router.Run(port)
}