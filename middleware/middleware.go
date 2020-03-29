package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckToken(c *gin.Context) {
	if c.GetHeader("token") == "LOZI" {
		c.Next()
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "token is invalid"})
		c.Abort()
		return
	}
}
