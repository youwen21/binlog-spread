package models

import (
	"time"
)

type RbacUserRole struct {
	UserRoleId     int       `gorm:"primary_key;" xorm:"not null pk autoincr INT(11)"`
	UserId         int       `xorm:"unique(user_role) INT(11)"`
	RoleId         int       `xorm:"unique(user_role) INT(11)"`
	AssignmentDate time.Time `xorm:"not null DATETIME"`
	IsDeleted      int       `xorm:"default 0 TINYINT(4)"`
	DeletedAt      time.Time `xorm:"DATETIME"`
}
