package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "12345")
		c.Next()
		log.SetPrefix("[Logger] ")
		latency := time.Since(t)
		status, method := c.Writer.Status(), c.Request.Method
		log.Println(method, status, latency)
	}
}
