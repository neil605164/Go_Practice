package router

import (
	"Go_Practice/app/handler"

	"github.com/gin-gonic/gin"
)

func LoadApiRouter(r *gin.Engine) {

	// 初始化 header 接口
	hand := handler.NewHandler()

	// 呼叫 API
	r.GET("/hand01", hand.MyHandler01)
	// 呼叫 Redis
	r.GET("/hand02", hand.MyHandler02)
	// 呼叫 DB 存資料
	r.POST("/hand03", hand.MyHandler03)
	// 呼叫 DB 取資料
	r.GET("/hand04", hand.MyHandler04)
}
