package command

import "github.com/urfave/cli"

var CmdNew = cli.Command{
	Name:   "new",
	Usage:  "provide the webserver which builds and serves the site",
	Action: runNew,
}

func runNew(c *cli.Context) error {

	return nil
}
