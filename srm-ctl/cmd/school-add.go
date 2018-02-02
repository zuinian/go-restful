package cmd

import (
	"fmt"
	"github.com/urfave/cli"

	"srm/srm-apiserver/client"
	"srm/srm-apiserver/models"
)

func init() {
	RegisterCommands(cli.Command{
		Name:        "school-add",
		Usage:       "",
		Description: "Registry school",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "address,a",
				Value: "china",
				Usage: "language for the greeting",
			},
			cli.StringFlag{
				Name:  "phone",
				Value: "10086",
				Usage: "language for the greeting",
			},
		},
		Action: SchoolAdd,
	})
}

func SchoolAdd(c *cli.Context) {
	fmt.Println(c.Args())

	srm := &client.SRMCClient{
		Host: "http://127.0.0.1:5000",
	}
	school := models.School{
		Name:    c.Args().Get(0),
		Address: c.String("address"),
		Phone: c.String("phone"),
	}
	err := srm.SchoolAdd(school)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ok")
}
