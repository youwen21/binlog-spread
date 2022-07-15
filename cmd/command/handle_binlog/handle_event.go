package handle_binlog

import (
	"fmt"
	"github.com/go-mysql-org/go-mysql/replication"
)


// HandleEvent mysql binlog event 处理
// https://www.cnblogs.com/f-zhao/p/7639888.html
// https://baijiahao.baidu.com/s?id=1649870303802210611&wfr=spider&for=pc
func HandleEvent(e *replication.BinlogEvent) {
	switch e.Header.EventType {
	// case replication.UNKNOWN_EVENT:
	// case replication.START_EVENT_V3:
	// 	return "StartEventV3"
	case replication.QUERY_EVENT: // 数据表结构变更和新增表，如 query:ALTER TABLE `codeper`.`user` ADD COLUMN `test` varchar(255) NULL AFTER `age`
		handleQueryEvent(e)
	// 	return "QueryEvent"
	// case replication.STOP_EVENT:
	// 	return "StopEvent"
	case replication.ROTATE_EVENT: // 上一个mysql binlog file结束标识
	// 	return "RotateEvent"
	// case replication.INTVAR_EVENT:
	// 	return "IntVarEvent"
	// case replication.LOAD_EVENT:
	// 	return "LoadEvent"
	// case replication.SLAVE_EVENT:
	// 	return "SlaveEvent"
	// case replication.CREATE_FILE_EVENT:
	// 	return "CreateFileEvent"
	// case replication.APPEND_BLOCK_EVENT:
	// 	return "AppendBlockEvent"
	// case replication.EXEC_LOAD_EVENT:
	// 	return "ExecLoadEvent"
	// case replication.DELETE_FILE_EVENT:
	// 	return "DeleteFileEvent"
	// case replication.NEW_LOAD_EVENT:
	// 	return "NewLoadEvent"
	// case replication.RAND_EVENT:
	// 	return "RandEvent"
	// case replication.USER_VAR_EVENT:
	// 	return "UserVarEvent"
	case replication.FORMAT_DESCRIPTION_EVENT: // 新mysql binlog file 开始标识
	// 	return "FormatDescriptionEvent"
	// case replication.XID_EVENT:
	// 	return "XIDEvent"
	// case replication.BEGIN_LOAD_QUERY_EVENT:
	// 	return "BeginLoadQueryEvent"
	// case replication.EXECUTE_LOAD_QUERY_EVENT:
	// 	return "ExectueLoadQueryEvent"
	case replication.TABLE_MAP_EVENT: // 有insert, update, delete)操作时，第来一条TABLE_MAP_EVENT 说明当下表结构
		// ev, _ := e.Event.(*replication.TableMapEvent)
		// e.Event.(replication.TableMapEvent).Table
		// FlushTableIdentifierNameMap(string(ev.Schema), string(ev.Table))
	// 	return "TableMapEvent"
	case replication.WRITE_ROWS_EVENTv0:
	// 	return "WriteRowsEventV0"
	case replication.UPDATE_ROWS_EVENTv0:
		fmt.Println("update rows event v0")
	// 	return "UpdateRowsEventV0"
	case replication.DELETE_ROWS_EVENTv0:
	// 	return "DeleteRowsEventV0"
	case replication.WRITE_ROWS_EVENTv1:
		handleWriteRowsEventV1(e)
	// 	return "WriteRowsEventV1"
	case replication.UPDATE_ROWS_EVENTv1:
		handleUpdateEventV1(e)
		//e.Dump(os.Stdout)
	// 	return "UpdateRowsEventV1"
	case replication.DELETE_ROWS_EVENTv1:
		// apputil.PrettyPrint(e)
		handleDeleteRowsEventV1(e)
	// 	return "DeleteRowsEventV1"
	// case replication.INCIDENT_EVENT:
	// 	return "IncidentEvent"
	// case replication.HEARTBEAT_EVENT:
	// 	return "HeartbeatEvent"
	// case replication.IGNORABLE_EVENT:
	// 	return "IgnorableEvent"
	// case replication.ROWS_QUERY_EVENT:
	// 	return "RowsQueryEvent"
	case replication.WRITE_ROWS_EVENTv2:
		handleWriteRowsEventV1(e)
	case replication.UPDATE_ROWS_EVENTv2: //
		fmt.Println("update rows event v2")
		handleUpdateEventV1(e)
	case replication.DELETE_ROWS_EVENTv2:
		handleDeleteRowsEventV1(e)
	case replication.GTID_EVENT: // 新GTID开始标识，它前面必然有一个PREVIOUS_GTIDS_EVENT事件
	// 	return "GTIDEvent"
	case replication.ANONYMOUS_GTID_EVENT:
	// 	return "AnonymousGTIDEvent"
	case replication.PREVIOUS_GTIDS_EVENT: //老GTID结束标识
	// 	return "PreviousGTIDsEvent"
	// case replication.MARIADB_ANNOTATE_ROWS_EVENT:
	// 	return "MariadbAnnotateRowsEvent"
	// case replication.MARIADB_BINLOG_CHECKPOINT_EVENT:
	// 	return "MariadbBinLogCheckPointEvent"
	// case replication.MARIADB_GTID_EVENT:
	// 	return "MariadbGTIDEvent"
	// case replication.MARIADB_GTID_LIST_EVENT:
	// 	return "MariadbGTIDListEvent"
	case replication.TRANSACTION_CONTEXT_EVENT:
	// 	return "TransactionContextEvent"
	// case replication.VIEW_CHANGE_EVENT:
	// 	return "ViewChangeEvent"
	// case replication.XA_PREPARE_LOG_EVENT:
	// 	return "XAPrepareLogEvent"

	default:
		return
	}
}

