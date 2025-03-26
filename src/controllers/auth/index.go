package auth

import (
	"app/src/services"
	"app/src/types"
	"app/src/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user types.User
	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user, err = services.Login(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	secret :=  os.Getenv("SECRET_KEY")
	token, err := utils.GenerateJwtToken(user.ID, secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}