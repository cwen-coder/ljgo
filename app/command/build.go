package command

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"git.cwengo.com/cwen/ljgo/app/lib"
	"github.com/urfave/cli"
)

var CmdBuild = cli.Command{
	Name:   "build",
	Usage:  "Generate blog to pubilc folder",
	Action: runBuild,
}

func runBuild(c *cli.Context) error {

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	build()

	go func() {
		<-signalChan
		fmt.Println()
		os.Exit(0)
	}()
	return nil
}

func build() {
	var article lib.Article
	err := article.ParseArticle("./source/article.md")
	if err != nil {
		log.Println(err)
	}
}
