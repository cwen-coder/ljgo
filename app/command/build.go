package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/qiniu/log"
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

	InitRootPath(c)
	build()

	go func() {
		<-signalChan
		fmt.Println()
		os.Exit(0)
	}()
	return nil
}

func build() {
	themePath := filepath.Join(rootPath, globalConfig.Site.Theme)
	partialPath := filepath.Join(themePath, "Tpl")
	files, err := filepath.Glob(filepath.Join(partialPath, "*.tpl"))
	if err != nil {
		log.Fatal(err)
	}
	for _, path := range files {
		_, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
	}
}
