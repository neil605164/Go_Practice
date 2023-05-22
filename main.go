package main

import (
	"Go_Practice/app/tool"
	"Go_Practice/internal/database"
	"fmt"
)

func init() {
	database.DBConn()
	database.DBPing()
	database.CheckTable()

	tool.SetupCasbin()
}

func main() {
	// r := gin.Default()
	// router.ProviderRouter(r)
	// r.Run(":3001")

	fmt.Println("Start")

	ok, err := tool.Enforcer.Enforce("neil", "it", "rd1", "read")
	if err != nil {
		fmt.Println("===>", err)
		return
	}

	if !ok {
		fmt.Println("access denied")
		return
	}

	fmt.Println("Pass")
}
