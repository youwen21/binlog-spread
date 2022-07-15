package models

import (
	"time"
)

type RbacPermission struct {
	PermissionId int       `gorm:"primary_key;" xorm:"not null pk autoincr INT(11)"`
	Method       string    `xorm:"CHAR(10)"`
	Source       string    `xorm:"VARCHAR(255)"`
	Title        string    `xorm:"not null default '' index CHAR(64)"`
	IsDeleted    int       `xorm:"default 0 TINYINT(4)"`
	DeletedAt    time.Time `xorm:"DATETIME"`
}
