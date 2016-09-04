package command

import (
	"fmt"

	"github.com/urfave/cli"
)

var CmdInfo = cli.Command{
	Name:   "info",
	Usage:  "provide ljgo info",
	Action: runInfo,
}

var (
	GitCommit string
	BuildTime string
)

func runInfo(c *cli.Context) error {
	if GitCommit == "" {
		GitCommit = "unknown"
	}
	if BuildTime == "" {
		BuildTime = "unknown"
	}
	fmt.Printf("git commit: %v\n", GitCommit)
	fmt.Printf("build time: %v\n", BuildTime)
	return nil
}
