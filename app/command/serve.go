package command

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"

	"git.cwengo.com/cwen/ljgo/app/config"

	"github.com/codegangsta/negroni"
	"github.com/qiniu/log"
	"github.com/urfave/cli"
	"gopkg.in/fsnotify.v1"
)

var CmdServe = cli.Command{
	Name:   "serve",
	Usage:  "provide the webserver which builds and serves the site",
	Action: runServe,
}

func runServe(c *cli.Context) error {
	cfg, err := config.New(c)
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	build(cfg)
	watch(c, cfg)
	serve(cfg)
	return nil
}

func serve(cfg *config.Config) {
	n := negroni.New()
	n.Use(negroni.NewStatic(http.Dir(cfg.PublicPath)))
	addr := cfg.Serve.Addr
	if addr == "" {
		addr = ":3000"
	}
	n.Run(addr)
}

func watch(c *cli.Context, cfg *config.Config) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op != fsnotify.Chmod {
					fmt.Println("modified file:", event.Name)
					runBuild(c)
				}
			case err := <-watcher.Errors:
				log.Errorf("error: %v", err)
			}
		}
	}()
	watcher.Add(filepath.Join(cfg.RootPath, "config.yml"))
	watchDir(watcher, filepath.Join(cfg.RootPath, "source"))
	watchDir(watcher, filepath.Join(cfg.RootPath, "themes"))
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
