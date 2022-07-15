package store

import (
	"binlog_spread/app/models"
	"binlog_spread/comps"
)

func SaveStateAbnormal(dbName string, tableName string, fieldName string, stateFrom string, stateTo string) {
	info := models.StateAbnormal{}
	info.DbName = dbName
	info.TableName = tableName
	info.FieldName = fieldName
	info.StateFrom = stateFrom
	info.StateTo = stateTo

	//gorm := mysql_conn.GetAssSession()
	gorm := comps.GetSession()
	gorm.Table("state_abnormal").Create(&info)
}
