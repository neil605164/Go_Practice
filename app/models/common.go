package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseIdModel struct {
	Id uint64 `gorm:"primaryKey;autoIncrement;unsigned"`
}

type BaseTimeFormat struct {
	CreatedAt time.Time      `gorm:"type:datetime comment '建立資料時間';autoCreateTime; NOT NULL"`
	UpdatedAt time.Time      `gorm:"type:datetime comment '資料最後更新時間';autoUpdateTime;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
