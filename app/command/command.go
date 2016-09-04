package command

import (
	"path/filepath"

	"github.com/qiniu/log"

	"github.com/urfave/cli"

	"git.cwengo.com/cwen/ljgo/app/library"
)

var globalConfig *library.Config
var rootPath string

func initConfig(c *cli.Context) {
	var err error
	if len(c.String("path")) > 0 {
		rootPath = c.String("path")
	} else {
		rootPath = "."
	}
	globalConfig, err = library.ParseConfig(filepath.Join(rootPath, "config.yml"))
	if err != nil {
		log.Fatalf("parse config.yml: %v", err)
	}
}
