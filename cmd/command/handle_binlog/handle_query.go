package handle_binlog

import (
	"binlog_spread/cmd/command/handle_binlog/api_binlog"
	"fmt"
	"github.com/go-mysql-org/go-mysql/replication"
)

func handleQueryEvent(e *replication.BinlogEvent) {
	ev, _ := e.Event.(*replication.QueryEvent)
	if string(ev.Schema) == "" {
		fmt.Println("event schema is empty")
		//e.Dump(os.Stdout)
		return
	}


	switch string(ev.Query){
	case "BEGIN", "begin":
		return
	default:
		api_binlog.BinlogQueryHandler.FlushDBTables(string(ev.Schema))
	}


}
