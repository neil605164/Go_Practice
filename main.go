package main

import (
	"Go_Practice/app/router"
	"Go_Practice/internal/database"

	"github.com/gin-gonic/gin"
)

func init() {
	database.DBConn()
	database.DBPing()
	database.CheckTable()
}

func main() {
	r := gin.Default()
	router.ProviderRouter(r)
	r.Run(":3001")
}
