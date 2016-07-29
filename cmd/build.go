package cmd

import "github.com/urfave/cli"

var CmdBuild = cli.Command{
	Name:  "build",
	Usage: "Generate blog to pubilc folder",
	Action: func(c *cli.Context) error {
	},
}
