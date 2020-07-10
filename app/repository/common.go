package repository

import "github.com/jinzhu/gorm"

// test 測試用
func testCreateRepository(db *gorm.DB) IDB {
	dbSingleton = &DB{
		dbcon: db,
	}

	return dbSingleton
}
