package command

import (
	"net/http"

	"github.com/codegangsta/negroni"
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
	initConfig(c)
	build(c)
	serve(c)
	return nil
}

func serve(c *cli.Context) {
	dir := rootPath + "/public"
	addr := c.String("addr")
	n := negroni.New()
	n.Use(negroni.NewStatic(http.Dir(dir)))
	if addr == "" {
		addr = ":3000"
	}
	n.Run(addr)
}

func watch() {

}
