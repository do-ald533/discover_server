package app

import (
	"github.com/urfave/cli"
)

func App() *cli.App {
	app := cli.NewApp()

	app.Name = "discover server"
	app.Usage = "Searches IPs and Names of servers"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "search for IPs in the internet",
			Flags: []cli.Flag{},
		},
	}

	return app
}
