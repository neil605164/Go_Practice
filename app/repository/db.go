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
		// CreatedAt: time.Now().UTC(),
		// UpdatedAt: time.Now().UTC(),
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

// 目前這情況，可能會有 test 連線覆蓋正常連線問題，需要在測試
func testCreateRepository(db *gorm.DB) IDB {
	// dbOnce.Do(func() {
	dbSingleton = &DB{
		dbcon: db,
	}
	// })
	return dbSingleton
}
