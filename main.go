package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 建立db連線
	// sql.Open("資料庫類型", "用戶名稱:用戶密碼@/資料庫名稱")
	db, err := sql.Open("mysql", "root:neil820724@/prac_golang")

	// 連線錯誤處理
	if err != nil {
		log.Fatal(err)
	}

	// 釋放連接池
	defer db.Close()

	//創建資料表
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS prac_golang.hello(world varchar(50))")
	if err != nil {
		log.Fatalln(err)
	}

	// 新增資料
	rs, err := db.Exec("INSERT INTO hello(world) VALUES ('hello world')")
	if err != nil {
		log.Fatalln(err)
	}

	// 檢查是否有存入成功
	rowCount, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("inserted %d rows", rowCount)

	// Query 查詢資料
	rows, err := db.Query("SELECT world FROM hello")
	if err != nil {
		log.Fatalln(err)
	}

	// 整理取出後的資料
	for rows.Next() {
		var s string
		// 檢查取資料時，是否有錯誤
		err = rows.Scan(&s)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("found row containing %q", s)
	}

	rows.Close()

	// 檢查rows.err()是否有錯誤
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	// QueryRow 查詢資料
	var str string
	err = db.QueryRow("SELECT world FROM hello LIMIT 1").Scan(&str)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("There is nit row")
		} else {
			log.Fatalln(err)
		}
	}

	log.Panicln("found a row", str)
}
