package command

import "github.com/urfave/cli"

var (
	pathFlag = cli.StringFlag{
		Name:  "path",
		Value: "template",
		Usage: "blog path",
	}

	addrFlag = cli.StringFlag{
		Name:  "addr",
		Value: "",
		Usage: "blog serve addr",
	}
)
