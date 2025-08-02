package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := c.Cookie("uid")
		if err != nil || uid == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
			return
		}
		c.Set("uid", uid)
		c.Next()
	}
}
