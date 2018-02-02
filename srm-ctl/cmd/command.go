package cmd

import "github.com/urfave/cli"

var MyCommands = []cli.Command{}

func RegisterCommands(command cli.Command) {
	MyCommands = append(MyCommands, command)
}
