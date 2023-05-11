package main

import (
	"Go_Practice/app/router"
	"Go_Practice/app/tool"
	"Go_Practice/internal/database"

	"github.com/gin-gonic/gin"
)

func init() {
	database.DBConn()
	database.DBPing()
	database.CheckTable()

	tool.SetupCasbin()
}

func main() {
	r := gin.Default()
	router.ProviderRouter(r)
	r.Run(":3001")
}
