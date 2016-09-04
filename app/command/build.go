package command

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"git.cwengo.com/cwen/ljgo/app/config"
	"git.cwengo.com/cwen/ljgo/app/library"
	"git.cwengo.com/cwen/ljgo/app/render"
	"git.cwengo.com/cwen/ljgo/app/util"
	"github.com/qiniu/log"
	"github.com/urfave/cli"
)

var CmdBuild = cli.Command{
	Name:  "build",
	Usage: "Generate blog to pubilc folder",
	Flags: []cli.Flag{
		pathFlag,
	},
	Action: runBuild,
}

func runBuild(c *cli.Context) error {
	cfg, err := config.New(c)
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	build(cfg)
	return nil
}

func build(cfg *config.Config) {
	partialPath := filepath.Join(cfg.ThemePath, "Tpl")
	partialTpl := buildPartialTpl(partialPath)
	articleTpl := buildTpl(filepath.Join(cfg.ThemePath, "article.html"), partialTpl, "article")
	indexTpl := buildTpl(filepath.Join(cfg.ThemePath, "index.html"), partialTpl, "index")
	aboutTpl := buildTpl(filepath.Join(cfg.ThemePath, "about.html"), partialTpl, "about")
	archiveTpl := buildTpl(filepath.Join(cfg.ThemePath, "archive.html"), partialTpl, "archive")
	tagTpl := buildTpl(filepath.Join(cfg.ThemePath, "tag.html"), partialTpl, "tag")

	cleanPatterns := []string{"static", "js", "css", "img", "vendor", "*.html", "*.xml"}
	cleanTpl(cfg.PublicPath, cleanPatterns)
	err := os.MkdirAll(cfg.PublicPath, 0777)
	if err != nil {
		log.Fatalf("create %v: %v", cfg.PublicPath, err)
	}

	articles := walkArticle(cfg.SourcePath)
	renderPage := render.New(cfg)
	renderPage.Index(indexTpl, articles)
	renderPage.Archive(archiveTpl, articles)
	renderPage.About(aboutTpl, filepath.Join(cfg.SourcePath, "about.md"))
	renderPage.Tags(tagTpl, articles)
	renderPage.RSS(articles)
	renderPage.Articles(articleTpl, articles)
	staticPath := filepath.Join(cfg.ThemePath, "static")
	copyStaticFile(staticPath, cfg.PublicPath)
}

func walkArticle(path string) library.Articles {
	articles := make(library.Articles, 0)
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		fileExt := strings.ToLower(filepath.Ext(path))
		if fileExt != ".md" {
			return nil
		}
		fileName := filepath.Base(path)
		noExtName := strings.TrimSuffix(fileName, ".md")
		if noExtName == "about" {
			return nil
		}
		var article library.Article
		err = article.ParseArticle(path)
		if err != nil {
			log.Errorf("parse article: %v", err)
			return nil
		}
		articles = append(articles, article)
		return nil
	})
	sort.Sort(articles)
	return articles
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
		htmlStr := "{{define \"" + fileName + "\"}}" + string(html) + "{{end}}"
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

func cleanTpl(cleanPath string, cleanPatterns []string) {
	for _, pattern := range cleanPatterns {
		files, err := filepath.Glob(filepath.Join(cleanPath, pattern))
		if err != nil {
			continue
		}
		for _, path := range files {
			os.RemoveAll(path)
		}
	}
}

func copyStaticFile(staticPath, publicPath string) {
	matches, err := filepath.Glob(staticPath)
	if err != nil {
		log.Fatalf("glob %v: %v", staticPath, err)
	}
	for _, srcPath := range matches {
		file, err := os.Stat(srcPath)
		if err != nil {
			log.Fatalf("copy static failed: %v", err)
		}
		filename := file.Name()
		destPath := filepath.Join(publicPath, filename)
		if file.IsDir() {
			err = util.CopyDir(srcPath, destPath)
			if err != nil {
				log.Fatalf("Copy %v: %v", filename, err)
			}
		} else {
			err = util.CopyFile(srcPath, destPath)
			if err != nil {
				log.Fatalf("Copy %v: %v", filename, err)
			}
		}
	}
}
