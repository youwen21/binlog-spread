package models

import (
	"time"
)

type RbacRolePermission struct {
	RolePermissionsId int       `gorm:"primary_key;" xorm:"not null pk autoincr INT(11)"`
	RoleId            int       `xorm:"not null unique(role_permission) INT(11)"`
	PermissionId      int       `xorm:"not null unique(role_permission) INT(11)"`
	AssignmentDate    time.Time `xorm:"not null DATETIME"`
	IsDeleted         int       `xorm:"default 0 TINYINT(4)"`
	DeletedAt         time.Time `xorm:"DATETIME"`
}
