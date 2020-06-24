package model

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// dbCon DB連線資料
type dbCon struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// masterPool 存放 db Master 連線池的全域變數
var masterPool *gorm.DB

// MasterConnect 建立 Master Pool 連線
func MasterConnect() (*gorm.DB, error) {
	var err error

	if masterPool != nil {
		return masterPool, nil
	}

	connString := composeString()
	masterPool, err = gorm.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	// 限制最大開啟的連線數
	masterPool.DB().SetMaxIdleConns(100)
	// 限制最大閒置連線數
	masterPool.DB().SetMaxOpenConns(2000)
	// 空閒連線 timeout 時間
	masterPool.DB().SetConnMaxLifetime(15 * time.Second)

	// 全局禁用表名复数
	// masterPool.SingularTable(true)
	// 開啟SQL Debug模式
	masterPool.LogMode(true)

	return masterPool, nil
}

// CheckTableIsExist 啟動main.go服務時，直接檢查所有 DB 的 Table 是否已經存在
func CheckTableIsExist() {
	db, err := MasterConnect()
	if err != nil {
		log.Fatalf("🔔🔔🔔 MASTER DB CONNECT ERROR: %v 🔔🔔🔔", err.Error())
	}

	// 會自動建置 DB Table
	err = db.Set("gorm:table_options", "comment '用戶資訊'").AutoMigrate(&User{}).Error
	if err != nil {
		panic(err)
	}
}

// composeString 組合DB連線前的字串資料
func composeString() string {
	db := dbCon{}

	db.Host = "localhost"
	db.Username = "root"
	db.Password = "root"
	db.Database = "go_prac"

	return fmt.Sprintf("%s:%s@tcp(%s:3307)/%s?timeout=5s&charset=utf8mb4&parseTime=True&loc=Local", db.Username, db.Password, db.Host, db.Database)
}
