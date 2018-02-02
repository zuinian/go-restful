package cmd

import (
	"github.com/urfave/cli"
	"fmt"
	"srm/srm-apiserver/client"
)

func init() {
	RegisterCommands(cli.Command{
		Name:        "school-students-get",
		Usage:       "school-students-get",
		Description: "school-students-get",
		Action:      schoolStudentsGet,
	})
}

func schoolStudentsGet(c *cli.Context) {
	srm := &client.SRMCClient{
		Host: "http://127.0.0.1:5000",
	}

	schoolName := c.Args().Get(0)
	res, err := srm.SchoolStudentsGet(schoolName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
