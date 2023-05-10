package models

type Role struct {
	BaseIdModel
	Name   string  `gorm:"type:varchar(255);not null;unique"`
	Admins []Admin `gorm:"many2many:admin_roles"`
	BaseTimeFormat
}
