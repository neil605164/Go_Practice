package middleware

import (
	"github.com/gin-gonic/gin"
)

func PermissionCheck(c *gin.Context) {
	c.Next()
}
