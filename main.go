package main

import (
	"binlog_spread/cmd/command"
	"binlog_spread/comps/sqlite3_conn"
	"binlog_spread/conf"
	"binlog_spread/lib"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

func main() {
	initProject()

	startProject()
}

func initProject() {
	if conf.Config.DbType == "sqlite" {
		if !lib.CheckFileIsExist(conf.Config.SqliteFile) {
			file, err := os.Create(conf.Config.SqliteFile)
			if err != nil {
				log.Fatal(err)
			}
			file.Close()

			sqlite3_conn.InitDatabase()

			time.Sleep(time.Second * 5)
		}
	}

}

func startProject() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "dev",
				Aliases: []string{"c"},
				Usage:   "开发运行调试",
				Action:  command.Dev,
			},
			{
				Name:  "binlog-start",
				Usage: "开始消费mysql binlog",

				Flags: []cli.Flag{
					&cli.StringFlag{Name: "config", Value: ""},
					&cli.StringFlag{Name: "host", Value: "127.0.0.1"},
					&cli.IntFlag{Name: "port", Value: 3306},
					&cli.StringFlag{Name: "username", Value: "root"},
					&cli.StringFlag{Name: "password", Value: ""},
					&cli.StringFlag{Name: "database", Value: ""},
					&cli.StringFlag{Name: "charset", Value: "utf8"},
					&cli.StringFlag{Name: "filter", Value: ".*\\..*"},
					&cli.IntFlag{Name: "server_id", Value: 0},
				},

				Action: command.BinlogClient.StartBinlogClient,
			},
			{
				Name:   "web-start",
				Usage:  "启动web服务器",
				Action: command.StartWebServer,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
