package main

import (
	"os"

	"github.com/urfave/cli"

	"srm/srm-ctl/cmd"
)

func main() {
	app := cli.NewApp()
	app.Name = "srm-ctl"
	app.Version = "1.0.0"
	app.Usage = "CLI for SRM!"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Usage: "path to config file",
		},
	}
	app.Commands = cmd.MyCommands
	app.Run(os.Args)
}