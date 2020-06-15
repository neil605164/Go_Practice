package router

import (
	"Go_Practice/app/handler"

	"github.com/gin-gonic/gin"
)

func LoadApiRouter(r *gin.Engine) {

	// 初始化 header 接口
	hand := handler.NewHandler()
	r.GET("/hand01", hand.MyHandler01)
}
