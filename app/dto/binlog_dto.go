package dto

import (
	"binlog_spread/app/models"
	"strconv"
)

type ApiBinlogForm struct {
	models.ApiBinlog
	PageParam
}

type ApiBinlogResult struct {
	Count int64              `json:"count"`
	List  []models.ApiBinlog `json:"list"`
}

type ApiBinlog models.ApiBinlog

type ApiBinlogText struct {
	models.ApiBinlog
	EventName string `json:"event_name"`
}

func GenApiBinlogText(binlog *models.ApiBinlog) ApiBinlogText {
	eventName := EventType(binlog.EventType).String()
	return ApiBinlogText{
		*binlog,
		eventName,
	}
}

//DiffRow 一行异同对比
type DiffRow struct {
	A ApiBinlogText `json:"a"`
	B ApiBinlogText `json:"b"`
}

// * 事件类型 *
type EventType int32

//以下代码： canal-go@v1.0.10/protocol/EntryProtocol.pb.go
const (
	EventType_EVENTTYPECOMPATIBLEPROTO2 EventType = 0
	EventType_INSERT                    EventType = 1
	EventType_UPDATE                    EventType = 2
	EventType_DELETE                    EventType = 3
	EventType_CREATE                    EventType = 4
	EventType_ALTER                     EventType = 5
	EventType_ERASE                     EventType = 6
	EventType_QUERY                     EventType = 7
	EventType_TRUNCATE                  EventType = 8
	EventType_RENAME                    EventType = 9
	// *CREATE INDEX*
	EventType_CINDEX EventType = 10
	EventType_DINDEX EventType = 11
	EventType_GTID   EventType = 12
	// * XA *
	EventType_XACOMMIT   EventType = 13
	EventType_XAROLLBACK EventType = 14
	// * MASTER HEARTBEAT *
	EventType_MHEARTBEAT EventType = 15
)

var EventType_name = map[int32]string{
	0:  "EVENTTYPECOMPATIBLEPROTO2",
	1:  "INSERT",
	2:  "UPDATE",
	3:  "DELETE",
	4:  "CREATE",
	5:  "ALTER",
	6:  "ERASE",
	7:  "QUERY",
	8:  "TRUNCATE",
	9:  "RENAME",
	10: "CINDEX",
	11: "DINDEX",
	12: "GTID",
	13: "XACOMMIT",
	14: "XAROLLBACK",
	15: "MHEARTBEAT",
}
var EventType_value = map[string]int32{
	"EVENTTYPECOMPATIBLEPROTO2": 0,
	"INSERT":                    1,
	"UPDATE":                    2,
	"DELETE":                    3,
	"CREATE":                    4,
	"ALTER":                     5,
	"ERASE":                     6,
	"QUERY":                     7,
	"TRUNCATE":                  8,
	"RENAME":                    9,
	"CINDEX":                    10,
	"DINDEX":                    11,
	"GTID":                      12,
	"XACOMMIT":                  13,
	"XAROLLBACK":                14,
	"MHEARTBEAT":                15,
}

func (x EventType) String() string {
	x32 := int32(x)
	s, ok := EventType_name[x32]
	if ok {
		return s
	}
	return strconv.Itoa(int(x32))
}
