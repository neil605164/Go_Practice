package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:neil820724@/homework-node")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln("err", err)
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var uid string
		var username string
		var password string
		var nickname string
		var user_state string
		var create_at string
		err = rows.Scan(&uid, &username, &password, &nickname, &user_state, &create_at)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("uid %q :", uid)
		log.Printf("username %q :", username)
		log.Printf("password %q :", password)
		log.Printf("nickname %q :", nickname)
		log.Printf("user_state %q :", user_state)
		log.Printf("ucreate_atid %q :", create_at)

	}
	rows.Close()

	// router := gin.Default()

	// router.GET("/", getting)
	// router.POST("/somePost", posting)

	// router.Run()
	fmt.Println("Hello World")
}

// func getting(c *gin.Context) {
// 	name := c.Query("test")
// 	fmt.Println(name)
// 	c.JSON(200, gin.H{
// 		"message": name,
// 	})
// }
