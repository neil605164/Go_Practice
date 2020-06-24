package main

import (
	"Go_Practice/app/model"
	"Go_Practice/router"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := gin.Default()

	// DB 連線
	_, err := model.MasterConnect()
	if err != nil {
		panic(err)
	}

	// DB Table 建立檢查
	model.CheckTableIsExist()

	// 載入 router
	router.LoadApiRouter(r)

	r.Run(":8484")
}
