package model

import "time"

// User 管理者帳號
type User struct {
	ID        int       `json:"id" gorm:"type:int(11) unsigned auto_increment comment '用戶ID';not null;primary_key"`
	Name      string    `json:"name" gorm:"column:name;type:varchar(30) comment '用戶名稱';not null;unique"`
	Phone     string    `json:"phone" gorm:"column:phone;type:varchar(30) comment '用戶電話';not null;unique"`
	Age       int       `json:"int" gorm:"column:int;type:varchar(30) comment '用戶電話';not null;unique"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:TIMESTAMP comment '資料建立時間'; default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:TIMESTAMP comment '資料最後更新時間';not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

// TableName 设置 User 的表名为 `user`
func (User) TableName() string {
	return "user"
}
