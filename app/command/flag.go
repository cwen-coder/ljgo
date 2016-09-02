package command

import "github.com/urfave/cli"

var (
	pathFlag = cli.StringFlag{
		Name:  "path",
		Value: "template",
		Usage: "blog path",
	}
)
