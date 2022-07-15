package common


func GetKey(dbName string, tableName string, fieldName string) string {
	return dbName + "_" + tableName + "_" + fieldName
}