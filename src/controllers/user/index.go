package album

import (
	"app/src/services"
	"app/src/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
    users := services.GetUsers()
    c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user types.User
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = services.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully!"})
}

func EditUser(c *gin.Context) {
	var user types.User
	err := c.Bind(&user)
	id := c.Param("id")
	user.ID = id
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = services.EditUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully!"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully!"})
}