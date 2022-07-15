package command

import (
	"binlog_spread/cmd/command/handle_binlog"
	"binlog_spread/cmd/command/handle_binlog/api_binlog"
	"binlog_spread/cmd/command/handle_binlog/state"
	"binlog_spread/comps/mysql_conn"
	"binlog_spread/comps/mysql_conn/mysqlutil"
	"binlog_spread/conf"
	"context"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
)

type binlogClient struct{}

var (
	BinlogClient = binlogClient{}
)

// mysql binlog 消费错误场景
// https://www.jianshu.com/p/ec4e626ae0b0

// MySQL二进制日志分析 - TABLE_MAP_EVENT
// https://www.cnblogs.com/little-star-2015/p/11736822.html

// 解析MySQL binlog --（4）TABLE_MAP_EVENT
// https://blog.51cto.com/yanzongshuai/2090758

// 解析MySQL binlog --（1）大致结构及event type
//https://blog.51cto.com/yanzongshuai/2085203

// https://blog.csdn.net/whatday/article/details/107918399
// golang string int int32 int64 float32 float64 time 互相转换

// Convert binary value as string to uint32 in Golang
// https://stackoverflow.com/questions/54814382/convert-binary-value-as-string-to-uint32-in-golang/54814575

//StartBinlogClient 消费mysql binlog
func (bc *binlogClient) StartBinlogClient(c *cli.Context) error {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(":2112", nil)
		if err != nil {
			panic(err)
		}
	}()

	// prepare
	api_binlog.Filter = conf.Config.BinlogMysql.Filter

	//初始化 - 数据-检查状态
	if conf.Config.EnableCheckState == "yes" {
		state.InitState()
	}

	//初始化 - binlog 数据流
	if conf.Config.EnableModelStream == "yes" {
		api_binlog.BinlogQueryHandler.InitDBTables()
	}

	// 获取binlog position
	// https://www.jianshu.com/p/f85ecae6e7df
	masterPosition, err := mysqlutil.GetMysqlPosition(mysql_conn.BinlogDb)
	if err != nil {
		log.Fatal(err)
	}

	// 生成binlog消费实例
	cfg := mysqlutil.GetBinlogSyncerConfig(conf.Config.BinlogMysql)
	syncer := replication.NewBinlogSyncer(cfg)
	u32 := uint32(masterPosition["Position"].(uint64))
	streamer, err := syncer.StartSync(mysql.Position{Name: masterPosition["File"].(string), Pos: u32})
	if err != nil {
		log.Fatal(err)
	}

	// 启动消费binlog
	for {
		ev, err := streamer.GetEvent(context.Background())
		if err != nil {
			log.Println("sync binlog error", err)
		}
		// Dump event
		//ev.Dump(os.Stdout)
		handle_binlog.HandleEvent(ev)
	}

	return nil
}
