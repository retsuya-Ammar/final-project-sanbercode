package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// middleware basic auth username admin and password admin
func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, pass, ok := c.Request.BasicAuth()
		if ok && user == "admin" && pass == "admin" {
			c.Next()
			return
		}
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
