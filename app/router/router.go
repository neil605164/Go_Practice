package router

import (
	"Go_Practice/app/handler"
	"Go_Practice/app/middleware"

	"github.com/gin-gonic/gin"
)

func ProviderRouter(r *gin.Engine) {
	r.GET("/", middleware.PermissionCheck, handler.GetAdmin)
}
