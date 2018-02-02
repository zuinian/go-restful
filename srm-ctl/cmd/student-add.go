package cmd

import (
	"fmt"
	"github.com/urfave/cli"

	"srm/srm-apiserver/client"
)

func init() {
	RegisterCommands(cli.Command{
		Name:        "student-add",
		Usage:       "",
		Description: "Registry school",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "gender",
				Value: "male",
				Usage: "",
			},
			cli.IntFlag{
				Name:  "age",
				Value: 18,
				Usage: "",
			},
			cli.StringFlag{
				Name:  "schoolName",
				Value: "xiada",
				Usage: "",
			},
		},
		Action: StudentAdd,
	})
}

func StudentAdd(c *cli.Context) {
	fmt.Println(c.Args())

	srm := &client.SRMCClient{
		Host: "http://127.0.0.1:5000",
	}
	studentReqBody := &client.StudentReqBody{
		Name:    c.Args().Get(0),
		Gender: c.String("gender"),
		Age: int8(c.Int("age")),
		SchoolName: c.String("schoolName"),
	}
	err := srm.StudentAdd(studentReqBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ok")
}
