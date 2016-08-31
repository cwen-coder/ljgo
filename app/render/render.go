package render

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/qiniu/log"

	"git.cwengo.com/cwen/ljgo/app/library"
	"git.cwengo.com/cwen/ljgo/app/util"
)

type Render struct {
	Site library.SiteConfig
	Path string
}

func New(site library.SiteConfig, path string) *Render {
	return &Render{
		Site: site,
		Path: path,
	}
}

func (r *Render) Articles(tpl template.Template, articles []library.Article) {
	for _, article := range articles {
		link := filepath.Join(r.Path, article.Link)
		outfile, err := os.Create(link)
		if err != nil {
			log.Fatalf("creat article %v: %v", link, err)
		}
		var data = make(map[string]interface{})
		data["Article"] = article
		data["Site"] = r.Site
		err = tpl.Execute(outfile, data)
		if err != nil {
			log.Fatalf("Execute %v: %v", link, err)
		}
		outfile.Close()
	}
}

func (r *Render) Index(tpl template.Template, articles []library.Article) {
	link := filepath.Join(r.Path, "index.html")
	outfile, err := os.Create(link)
	defer outfile.Close()
	if err != nil {
		log.Fatalf("creat index.html: %v", err)
	}
	var data = make(map[string]interface{})
	data["Articles"] = articles
	data["Site"] = r.Site

	err = tpl.Execute(outfile, data)
	if err != nil {
		log.Fatalf("Execute index.html: %v", err)
	}
}

func (r *Render) Archive(tpl template.Template, articles []library.Article) {
	link := filepath.Join(r.Path, "archive.html")
	outfile, err := os.Create(link)
	defer outfile.Close()
	if err != nil {
		log.Fatalf("creat archive.html: %v", err)
	}
	var data = make(map[string]interface{})
	data["Articles"] = articles
	data["Site"] = r.Site

	err = tpl.Execute(outfile, data)
	if err != nil {
		log.Fatalf("Execute archive.html: %v", err)
	}
}

func (r *Render) About(tpl template.Template, path string) {
	if _, err := os.Stat(path); err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("open about: %v", err)
		}
		file, err := os.Create(path)
		if err != nil {
			log.Fatalf("create about: %v", err)
		}
		file.Close()
	}
	d, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("read about: %v", err)
	}
	aboutData := util.ParseMarkdown(string(d))

	link := filepath.Join(r.Path, "about.html")
	outfile, err := os.Create(link)
	defer outfile.Close()
	if err != nil {
		log.Fatalf("creat about.html: %v", err)
	}
	var data = make(map[string]interface{})
	data["About"] = aboutData
	data["Site"] = r.Site

	err = tpl.Execute(outfile, data)
	if err != nil {
		log.Fatalf("Execute about.html: %v", err)
	}
}
