package command

import (
	"log"

	"git.cwengo.com/cwen/ljgo/app/lib"
)

var globalConfig *lib.Config

func init() {
	var err error
	globalConfig, err = lib.ParseConfig("./themes/config.yml")
	if err != nil {
		log.Fatal(err)
	}
}
