package command

import (
	"log"

	"github.com/urfave/cli"

	"git.cwengo.com/cwen/ljgo/app/library"
)

var globalConfig *library.Config
var rootPath string

func init() {
	var err error
	globalConfig, err = library.ParseConfig("./themes/config.yml")
	if err != nil {
		log.Fatal(err)
	}
}

func InitRootPath(c *cli.Context) {
	if len(c.Args()) > 0 {
		rootPath = c.Args()[0]
	} else {
		rootPath = "."
	}
}

//type Collections []interface{}

//func (c Collections) Len() int {
//return len(c)
//}

//func (c Collections) Swap(i, j int) {
//c[i], c[j] = c[i], c[j]
//}

//func (c Collections) Less(i, j int) bool {

//}
