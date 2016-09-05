package command

import "github.com/urfave/cli"

var (
	addrFlag = cli.StringFlag{
		Name:  "addr",
		Value: "",
		Usage: "blog serve addr",
	}
)
