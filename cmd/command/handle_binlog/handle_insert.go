package handle_binlog

import (
	"binlog_spread/app/models"
	"binlog_spread/cmd/command/handle_binlog/api_binlog"
	"binlog_spread/cmd/command/handle_binlog/store"
	"binlog_spread/conf"
	"encoding/json"
	"fmt"
	"github.com/go-mysql-org/go-mysql/replication"
	"strings"
)

func handleWriteRowsEventV1(e *replication.BinlogEvent) {
	ev, _ := e.Event.(*replication.RowsEvent)

	if conf.Config.EnableModelStream == "yes" {
		insertRoutineModelStream(ev)
	}
}

func insertRoutineModelStream(ev *replication.RowsEvent) {
	dbName := string(ev.Table.Schema)
	tableName := string(ev.Table.Table)
	ok := api_binlog.MatchTable(dbName, tableName)
	if !ok {
		fmt.Println("skip write", dbName, ".", tableName)
		return
	}

	var streams []models.ApiBinlog2

	stream := &models.ApiBinlog2{}
	stream.DbName = dbName
	stream.TableName = tableName
	stream.TransactionTag = ""
	stream.EventType = 1 // 此处是canal定义，和原mysql binlog event type 不同

	for i := 0; i < len(ev.Rows); i++ {
		var allColumns []string
		var updatedColumns []string
		updatedData := make(map[string]interface{})

		tableSchema := api_binlog.DBTables[string(ev.Table.Schema)+"."+string(ev.Table.Table)]
		for idx, value := range ev.Rows[i] {
			allColumns = append(allColumns, tableSchema[idx])
			updatedColumns = append(updatedColumns, tableSchema[idx])
			updatedData[tableSchema[idx]] = value
		}

		stream.Columns = strings.Join(allColumns, ",")
		stream.UpdateColumns = strings.Join(updatedColumns, ",")

		b, _ := json.Marshal(updatedData)
		stream.UpdateValue = string(b)

		streams = append(streams, *stream)
	}

	store.StreamAddRows(streams)
}
