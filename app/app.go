package app

import (
	"github.com/urfave/cli"
)

func App() *cli.App {
	app := cli.NewApp()

	app.Name = "discover server"
	app.Usage = "Searches IPs and Names of servers"

	return app
}
