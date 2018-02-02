package cmd

import (
	"github.com/urfave/cli"
	"fmt"
	"srm/srm-apiserver/client"
)

func init() {
	RegisterCommands(cli.Command{
		Name:        "schools-get",
		Usage:       "schools-get",
		Description: "schools-get",
		Action:      schoolsGet,
	})
}

func schoolsGet(c *cli.Context) {
	srm := &client.SRMCClient{
		Host: "http://127.0.0.1:5000",
	}

	res, err := srm.SchoolsGet()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
