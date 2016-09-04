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
		Value: "localhost:3000",
		Usage: "blog serve addr",
	}
)
