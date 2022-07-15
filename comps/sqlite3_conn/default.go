package sqlite3_conn

import (
	"binlog_spread/conf"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//var once sync.Once

var (
	SpreadLiteDb *gorm.DB
)

func init() {
	if conf.Config.DbType == "sqlite" {
		db, err := gorm.Open(sqlite.Open(conf.Config.SqliteFile), &gorm.Config{})
		fmt.Println(err)
		if err != nil {
			panic(err)
		}
		SpreadLiteDb = db
	}
}

func InitDatabase() {
	db, err := SpreadLiteDb.DB()
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(tableSchema)
	if err != nil {
		panic(err)
	}
}
