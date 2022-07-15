package api_binlog

import (
	"binlog_spread/comps/mysql_conn"
	"binlog_spread/comps/mysql_conn/mysqlutil"
	"fmt"
	"regexp"
	"strings"
)

type binlogQueryHandler struct {
}

//TableColumnIdentify 表字段ID对应字段名
type TableColumnIdentify map[int]string

//DBTable  数据库.数据表map类型
type DBTable map[string]TableColumnIdentify

var (
	//Filter 定义过滤
	Filter string

	// DBTables 数据库.数据表map配置实例
	DBTables DBTable

	BinlogQueryHandler binlogQueryHandler
)

func init() {
	DBTables = make(DBTable, 0)
}

func (b binlogQueryHandler) InitDBTables() {
	dbs, _ := mysqlutil.GetDatabases(mysql_conn.BinlogDb.Debug())

	for _, v := range dbs {
		if v == "" {
			continue
		}
		b.FlushDBTables(v)
	}
}

func (b binlogQueryHandler) FlushDBTables(dbName string) {
	tables, _ := mysqlutil.GetTableNames(mysql_conn.BinlogDb.Debug(), dbName)
	for _, tableName := range tables {
		b.FlushTableIdentifierNameMap(dbName, tableName)
	}
}

// https://dev.mysql.com/doc/refman/8.0/en/identifiers.html
func (b binlogQueryHandler) FlushTableIdentifierNameMap(db string, table string) {
	if !MatchTable(db, table) {
		return
	}

	sql := "show full columns from `" + table + "` from `" + db + "`"
	// sql := "show full FIELDS from `user` from `codeper`"

	var results []map[string]interface{}

	err := mysql_conn.BinlogDb.Debug().Raw(sql).Find(&results).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	columns := TableColumnIdentify{}
	for i, v := range results {
		columns[i] = v["Field"].(string)
	}

	DBTables[db+"."+table] = columns
}

//MatchTable 检查table 是否在正则中
func MatchTable(db string, table string) bool {
	if true == SkipDb(db) {
		return false
	}

	filters := strings.Split(Filter, ",")
	for _, v := range filters {
		ok, _ := regexp.Match(v, []byte(db+"."+table))
		if ok {
			return true
		}
	}

	return false
}

func SkipDb(db string) bool {
	if db == "mysql" {
		return true
	}

	return false
}
