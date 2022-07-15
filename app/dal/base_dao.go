package dal

import (
	"binlog_spread/comps"
	"gorm.io/gorm"
)

type baseDAO struct{

}

func (d *baseDAO) newEngine() *gorm.DB {
	return comps.GetDb()
}

func (d *baseDAO) newSession() *gorm.DB {
	return comps.GetSession()
}