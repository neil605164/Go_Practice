package models

type Admin struct {
	BaseIdModel
	Username string `gorm:"type:varchar(255);not null;unique;comment:帳號"`
	Name     string `gorm:"type:varchar(255);not null;comment:名稱"`
	BaseTimeFormat
	Roles []Role `gorm:"many2many:admin_roles"`
}

type AdminRole struct {
	AdminId uint64 `gorm:"index:admin_id"`
	RoleId  uint64 `gorm:"index:role_id"`
}
