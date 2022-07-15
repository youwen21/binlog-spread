package comps

import (
	"binlog_spread/comps/mysql_conn"
	"binlog_spread/comps/sqlite3_conn"
	"binlog_spread/conf"
	"gorm.io/gorm"
)

func GetDb() *gorm.DB {
	if conf.Config.DbType == "sqlite"{
		return sqlite3_conn.SpreadLiteDb
	}

	if conf.Config.DbType == "mysql"{
		return mysql_conn.SpreadDb
	}

	return nil
}

func GetSession() *gorm.DB {
	if conf.Config.DbType == "sqlite"{
		return sqlite3_conn.SpreadLiteDb.Session(&gorm.Session{})
	}

	if conf.Config.DbType == "mysql"{
		return mysql_conn.SpreadDb.Session(&gorm.Session{})
	}

	return nil
}