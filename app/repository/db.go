package repository

import (
	"Go_Practice/app/global/structs"
	"Go_Practice/app/model"
	"sync"

	"github.com/jinzhu/gorm"
)

// IDB DB 接口
type IDB interface {
	SetUserInfo(req structs.RawData) (err error)
	GetUserInfo() (resp []model.User, err error)
	UpdateUserInfo(reqMap map[string]interface{}) (err error)
	DeleteUserInfo(id int) (err error)
}

// DB 存取值
type DB struct {
	dbcon *gorm.DB
}

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

	if db.dbcon == nil {
		// 取 DB 連線
		db.dbcon, err = model.MasterConnect()
		if err != nil {
			return
		}
	}

	// new user
	user := model.User{
		Name:  req.Name,
		Phone: req.Phone,
		Age:   req.Age,
	}

	if err = db.dbcon.Create(&user).Error; err != nil {
		return
	}

	return
}

// GetUserInfo 取用戶資訊
func (db *DB) GetUserInfo() (resp []model.User, err error) {

	// 取 DB 連線
	if db.dbcon == nil {
		// 取 DB 連線
		db.dbcon, err = model.MasterConnect()
		if err != nil {
			return
		}
	}

	if err = db.dbcon.Find(&resp).Error; err != nil {
		return
	}
	return
}

// UpdateUserInfo 更新用戶資訊
func (db *DB) UpdateUserInfo(reqMap map[string]interface{}) (err error) {
	if db.dbcon == nil {
		// 取 DB 連線
		db.dbcon, err = model.MasterConnect()
		if err != nil {
			return
		}
	}

	// update user
	return db.dbcon.Model(&model.User{}).Updates(reqMap).Error
}

// DeleteUserInfo 刪除用戶資料
func (db *DB) DeleteUserInfo(id int) (err error) {
	if db.dbcon == nil {
		// 取 DB 連線
		db.dbcon, err = model.MasterConnect()
		if err != nil {
			return
		}
	}

	// delete user
	return db.dbcon.Delete(&model.User{ID: id}).Error
}
