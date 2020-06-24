package repository

import (
	"Go_Practice/app/global/structs"
	"Go_Practice/app/model"
	"sync"
)

// IDB DB 接口
type IDB interface {
	SetUserInfo(req structs.RawData) (err error)
}

// DB 存取值
type DB struct{}

var dbSingleton *DB
var dbOnce sync.Once

// DBIns 獲得單例對象
func DBIns() IDB {
	dbOnce.Do(func() {
		dbSingleton = &DB{}
	})
	return dbSingleton
}

// SetUserInfo 寫入用戶資訊
func (db *DB) SetUserInfo(req structs.RawData) (err error) {
	// 取 DB 連線
	dbcon, err := model.MasterConnect()
	if err != nil {
		return
	}

	// new user
	user := model.User{
		Name:  req.Name,
		Phone: req.Phone,
		Age:   req.Age,
	}

	if err = dbcon.Create(&user).Error; err != nil {
		return
	}

	return
}
