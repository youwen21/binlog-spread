package command

import (
	"binlog_spread/router"
	"github.com/urfave/cli/v2"
)

func StartWebServer(c *cli.Context) error {
	router.StartHttpServer()
	return nil
}
