package command

import (
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"

	"github.com/codegangsta/negroni"
	"github.com/qiniu/log"
	"github.com/urfave/cli"
	"gopkg.in/fsnotify.v1"
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
	build()
	watch(c)
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

func watch(c *cli.Context) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	// defer watcher.Close()

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Info(event)
				if event.Op != fsnotify.Chmod {
					log.Println("modified file:", event.Name)
					runBuild(c)
				}
			case err := <-watcher.Errors:
				log.Errorf("error: %v", err)
			}
		}
	}()
	watcher.Add(filepath.Join(rootPath, "config.yml"))
	watchDir(watcher, filepath.Join(rootPath, "source"))
	watchDir(watcher, filepath.Join(rootPath, "themes"))
}

func watchDir(watcher *fsnotify.Watcher, srcDir string) {
	watcher.Add(srcDir)
	dir, _ := ioutil.ReadDir(srcDir)
	for _, d := range dir {
		if d.IsDir() {
			watchDir(watcher, path.Join(srcDir, d.Name()))
		}
	}
}
