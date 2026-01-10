package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, pass, ok := c.Request.BasicAuth()
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		adminUser := os.Getenv("Username")
		adminPass := os.Getenv("Password")

		if user != adminUser || pass != adminPass {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid credential",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
