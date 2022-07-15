package models

import (
	"time"
)

type RbacRole struct {
	RoleId      int       `gorm:"primary_key;" xorm:"not null pk autoincr INT(11)"`
	Title       string    `xorm:"not null index VARCHAR(128)"`
	Description string    `xorm:"not null TEXT"`
	IsDeleted   int       `xorm:"default 0 TINYINT(4)"`
	DeletedAt   time.Time `xorm:"DATETIME"`
}
