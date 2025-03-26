package middlewares

import (
	"app/src/types"
	"app/src/utils"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func PublicMiddleware(paths []types.Public) gin.HandlerFunc {
	return func(c *gin.Context) {
		isPublic := false
		for _, path := range paths {
			if path.Path == c.Request.URL.Path && path.Method == c.Request.Method {
				isPublic = true
				break
			}
		}
		c.Set("isPublic", isPublic)
		c.Next()
	}
}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isPublic := c.MustGet("isPublic").(bool)
		if isPublic {
			c.Next()
			return
		}
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token = strings.Split(token, " ")[1]
		claims, err := utils.ValidateJwtToken(token, os.Getenv("SECRET_KEY"))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userId", claims.ID)
		c.Next()
	}
}