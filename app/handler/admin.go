package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAdmin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "success")
}
