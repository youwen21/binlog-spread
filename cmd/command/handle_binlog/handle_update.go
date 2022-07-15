package handle_binlog

import (
	"binlog_spread/app/models"
	"binlog_spread/cmd/command/handle_binlog/api_binlog"
	"binlog_spread/cmd/command/handle_binlog/state"
	"binlog_spread/cmd/command/handle_binlog/store"
	"binlog_spread/conf"
	lib2 "binlog_spread/lib"
	"encoding/json"
	"fmt"
	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/google/go-cmp/cmp"
	"strings"
)

func handleUpdateEventV1(e *replication.BinlogEvent) {
	ev, _ := e.Event.(*replication.RowsEvent)

	if conf.Config.EnableCheckState == "yes" {
		go updateRoutineStatRule(ev)
	}

	if conf.Config.EnableModelStream == "yes" {
		updateRoutineModelStream(ev)
	}
}

func updateRoutineModelStream(ev *replication.RowsEvent) {
	dbName := string(ev.Table.Schema)
	tableName := string(ev.Table.Table)
	ok := api_binlog.MatchTable(dbName, tableName)
	if !ok {
		fmt.Println("skip update", dbName, ".", tableName)
		return
	}
	tableSchema := api_binlog.DBTables[dbName+"."+tableName]

	var streams []models.ApiBinlog2
	stream := models.ApiBinlog2{}
	stream.DbName = dbName
	stream.TableName = tableName
	stream.TransactionTag = ""
	stream.EventType = 2 // 此处是canal定义，和原mysql binlog event type 不同

	for i := 0; i < len(ev.Rows); i = i + 2 {
		var allColumns []string
		var updatedColumns []string
		updatedData := make(map[string]interface{})
		next := i + 1
		for idx, value := range ev.Rows[i] {
			fieldName := tableSchema[idx]
			allColumns = append(allColumns, fieldName)

			// go类型断言
			// https://www.jianshu.com/p/787cf3a41ccb
			// mysql 反回字段interface类型， 获取value参考
			// /Users/owen/go/pkg/mod/github.com/go-xorm/xorm@v0.7.9/session_query.go
			if !cmp.Equal(value, ev.Rows[next][idx]) {
				updatedColumns = append(updatedColumns, fieldName)
				//strValue := fmt.Sprintf("%s", ev.Rows[next][idx])
				//strValue := common.GetValueString(ev.Rows[next][idx])
				strValue, _ := lib2.GetStringValue(ev.Rows[next][idx])
				updatedData[fieldName] = strValue
			}
		}

		stream.Columns = strings.Join(allColumns, ",")
		stream.UpdateColumns = strings.Join(updatedColumns, ",")

		b, _ := json.Marshal(updatedData)
		stream.UpdateValue = string(b)

		streams = append(streams, stream)
	}

	store.StreamAddRows(streams)
}

func updateRoutineStatRule(ev *replication.RowsEvent) {
	dbName := string(ev.Table.Schema)
	tableName := string(ev.Table.Table)
	tableSchema := api_binlog.DBTables[dbName+"."+tableName]

	for i := 0; i < len(ev.Rows); i = i + 2 {
		next := i + 1
		for idx, value := range ev.Rows[i] {
			if cmp.Equal(value, ev.Rows[next][idx]) {
				continue
			}

			fieldName := tableSchema[idx]
			// 校验状态流
			classId, err := state.GetStatClassId(dbName, tableName, fieldName)
			if nil != err {
				//fmt.Println(dbName, tableName, fieldName, err)
				continue
			}

			from, _ := lib2.GetStringValue(value)
			to, _ := lib2.GetStringValue(ev.Rows[next][idx])
			check, err := state.CheckStatDirection(classId, from, to)
			// 流程变更不合规， 做一些通知URL, 钉钉，记录库等
			if !check {
				fmt.Println(dbName, tableName, fieldName, "classId:", classId, "from:", from, "to:", to, err)
				store.SaveStateAbnormal(dbName, tableName, fieldName, from, to)
			}
		}
	}
}
