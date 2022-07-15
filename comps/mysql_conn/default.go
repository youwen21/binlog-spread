package mysql_conn

import (
	"binlog_spread/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	SpreadDb *gorm.DB
	BinlogDb *gorm.DB
)

func init() {
	if conf.Config.DbType == "mysql" {
		initSpreadDb()
	}

	initBinlogDb()
}

func initSpreadDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=Local&parseTime=true",
		conf.Config.DefaultMysql.Username,
		conf.Config.DefaultMysql.Password,
		conf.Config.DefaultMysql.Host,
		conf.Config.DefaultMysql.Port,
		conf.Config.DefaultMysql.Database,
		conf.Config.DefaultMysql.Charset,
	)
	mysqlDb := mysql.Open(dsn)

	db, err := gorm.Open(mysqlDb, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(1)
	sqlDb.SetMaxOpenConns(50)

	SpreadDb = db
}

func initBinlogDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=Local&parseTime=true",
		conf.Config.BinlogMysql.Username,
		conf.Config.BinlogMysql.Password,
		conf.Config.BinlogMysql.Host,
		conf.Config.BinlogMysql.Port,
		conf.Config.BinlogMysql.Database,
		conf.Config.BinlogMysql.Charset,
	)
	mysqlDb := mysql.Open(dsn)

	db, err := gorm.Open(mysqlDb, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(1)
	sqlDb.SetMaxOpenConns(50)

	BinlogDb = db
}
