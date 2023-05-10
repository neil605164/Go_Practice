package database

import (
	"Go_Practice/app/models"
	"fmt"

	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var Db *gorm.DB

func DBConn() (*gorm.DB, error) {

	var err error
	var option = "%s:%s@tcp(%s:%s)/%s?loc=Local&parseTime=True"

	dsn_master := fmt.Sprintf(option,
		"root",
		"root",
		"127.0.0.1",
		"3306",
		"order",
	)

	dsn_slave := fmt.Sprintf(option,
		"root",
		"root",
		"127.0.0.1",
		"3306",
		"order",
	)

	// 連接gorm
	Db, err = gorm.Open(mysql.Open(dsn_master), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	Db.Use(
		dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(dsn_master)},
			Replicas: []gorm.Dialector{mysql.Open(dsn_slave)},
			Policy:   dbresolver.RandomPolicy{}}).
			// 空閒連線 timeout 時間
			SetConnMaxIdleTime(15 * time.Second).
			// 空閒連線 timeout 時間
			SetConnMaxLifetime(15 * time.Second).
			// 限制最大閒置連線數
			SetMaxIdleConns(100).
			// 限制最大開啟的連線數
			SetMaxOpenConns(2000),
	)

	sqlDB, err := Db.DB()
	if err != nil {
		return nil, err
	}

	// 限制最大閒置連線數
	sqlDB.SetMaxIdleConns(100)
	// 限制最大開啟的連線數
	sqlDB.SetMaxOpenConns(2000)
	// 空閒連線 timeout 時間
	sqlDB.SetConnMaxLifetime(15 * time.Second)

	return Db, nil
}

func DBPing() {
	dbcon, err := Db.DB()
	if err != nil {
		log.Fatalf("🔔🔔🔔 CONNECT MASTER DB ERROR: %v 🔔🔔🔔", err.Error())
	}

	err = dbcon.Ping()
	if err != nil {
		log.Fatalf("🔔🔔🔔 PING MASTER DB ERROR: %v 🔔🔔🔔", err.Error())
	}
}

// CheckTable 啟動main.go服務時，直接檢查所有 DB 的 Table 是否已經存在
func CheckTable() {

	// 會自動建置 DB Table
	err := Db.Set("gorm:table_options", "comment '後台帳戶管理'").AutoMigrate(&models.Admin{})
	if err != nil {
		panic(err.Error())
	}

	err = Db.Set("gorm:table_options", "comment '後台帳戶管理與角色關聯表'").AutoMigrate(&models.AdminRole{})
	if err != nil {
		panic(err.Error())
	}

	err = Db.Set("gorm:table_options", "comment '角色'").AutoMigrate(&models.Role{})
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		log.Fatalf("🔔🔔🔔 MIGRATE MASTER TABLE ERROR: %v 🔔🔔🔔", err.Error())
	}

}
