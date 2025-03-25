package route

import (
	AlbumRouter "app/src/routes/user"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine, port string) {
	AlbumRouter.UserRouter(router)
	router.Run(port)
}