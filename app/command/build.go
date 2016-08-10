package command

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
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
	partialTpl := buildPartialTpl(partialPath)

	articleTpl := buildTpl(filepath.Join(themePath, "article.html"), partialTpl, "article")
	log.Info(articleTpl)
}

func buildPartialTpl(path string) string {
	files, err := filepath.Glob(filepath.Join(path, "*.tpl"))
	if err != nil {
		log.Fatal(err)
	}
	var partialTpl string
	for _, path := range files {
		html, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		fileName := filepath.Base(path)
		fileName = strings.TrimSuffix(strings.TrimPrefix(fileName, "T."), ".tpl")
		htmlStr := "{define \"" + fileName + " \"}" + string(html) + "{end}"
		partialTpl += htmlStr
	}
	return partialTpl
}

func buildTpl(path string, partialTpl string, name string) template.Template {
	html, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	htmlStr := string(html) + partialTpl

	tpl, err := template.New(name).Parse(htmlStr)
	if err != nil {
		log.Fatal(err)
	}
	return *tpl
}
