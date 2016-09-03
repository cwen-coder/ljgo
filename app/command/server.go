package command

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli"
)

var CmdServer = cli.Command{
	Name:  "serve",
	Usage: "provide the webserver which builds and serves the site",
	Flags: []cli.Flag{
		pathFlag,
		addrFlag,
	},
	Action: runServe,
}

func runServe(c *cli.Context) error {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	Init(c)
	build()
	serve()
	go func() {
		<-signalChan
		fmt.Println()
		os.Exit(0)
	}()
	return nil
}

func serve() {
}

func watch() {
}
