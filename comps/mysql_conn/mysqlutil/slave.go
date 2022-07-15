package mysqlutil

import (
	"binlog_spread/conf"
	"github.com/go-mysql-org/go-mysql/replication"
	"gorm.io/gorm"
)

func GetMysqlPosition(db *gorm.DB) (map[string]interface{}, error) {
	sql := "show master status"
	results := make(map[string]interface{})

	err := db.Raw(sql).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func GetBinlogSyncerConfig(conf conf.BinlogMysql) replication.BinlogSyncerConfig {
	cfg := replication.BinlogSyncerConfig{
		ServerID: uint32(conf.ServerId),
		Flavor:   "mysql",
		Host:     conf.Host,
		Port:     uint16(conf.Port),
		User:     conf.Username,
		Password: conf.Password,
	}

	return cfg
}

func GetDatabases(gorm *gorm.DB) ([]string, error) {
	sql := "show databases"

	var result []map[string]interface{}
	err := gorm.Raw(sql).Find(&result).Error
	if err != nil {
		return nil, err
	}

	var dbs []string
	for _, v := range result {
		if v["Database"] == "mysql" || v["Database"] == "information_schema" || v["Database"] == "performance_schema" || v["Database"] == "sys" {
			continue
		}
		dbs = append(dbs, v["Database"].(string))
	}

	return dbs, nil
}

func GetTableNames(gorm *gorm.DB, db string) ([]string, error) {
	sql := "show tables from `" + db + "`"

	var result []map[string]interface{}
	err := gorm.Raw(sql).Find(&result).Error
	if err != nil {
		return nil, err
	}

	var tables []string

	for _, v := range result {
		tables = append(tables, v["Tables_in_"+db].(string))
	}

	return tables, nil
}
