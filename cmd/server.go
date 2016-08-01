package cmd

import "github.com/urfave/cli"

var CmdServer = cli.Command{
	Name:  "serve",
	Usage: "provide the webserver which builds and serves the site",
	Action: func(c *cli.Context) error {
		return nil
	},
}
