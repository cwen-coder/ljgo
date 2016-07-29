package cmd

import "github.com/urfave/cli"

var CmdPublish = cli.Command{
	Name:  "publish",
	Usage: "Generate blog to pubilc folder and publish",
	Action: func(c *cli.Context) error {
	},
}
