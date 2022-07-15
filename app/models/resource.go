package models

import (
	"time"
)

type Resource struct {
	ResourceId    int       `gorm:"primary_key;" xorm:"not null pk autoincr INT(11)"`
	ResourceGroup string    `xorm:"comment('分组') VARCHAR(20)"`
	ResourceName  string    `xorm:"VARCHAR(255)"`
	Identity      string    `xorm:"VARCHAR(255)"`
	IsDeleted     int       `xorm:"default 0 TINYINT(4)"`
	DeletedAt     time.Time `xorm:"DATETIME"`
}
