package render

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

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

func (r *Render) Articles(tpl template.Template, articles library.Articles) {
	for _, article := range articles {
		link := filepath.Join(r.Path, article.Link)
		outfile, err := os.Create(link)
		if err != nil {
			log.Fatalf("creat article %v: %v", link, err)
		}
		var data = make(map[string]interface{})
		data["Article"] = article
		data["Site"] = r.Site
		data["Title"] = article.ConfigArticle.Title
		err = tpl.Execute(outfile, data)
		if err != nil {
			log.Fatalf("Execute %v: %v", link, err)
		}
		outfile.Close()
	}
}

func (r *Render) Index(tpl template.Template, articles library.Articles) {
	total := len(articles)
	indexCount := total / r.Site.Limit
	rest := total % r.Site.Limit
	if rest != 0 {
		indexCount++
	}

	for i := 0; i < indexCount; i++ {
		prev := "index" + strconv.Itoa(i) + ".html"
		next := "index" + strconv.Itoa(i+2) + ".html"
		link := filepath.Join(r.Path, "index"+strconv.Itoa(i+1)+".html")
		if i == 0 {
			link = filepath.Join(r.Path, "index.html")
			prev = ""
		}
		outfile, err := os.Create(link)
		defer outfile.Close()
		if err != nil {
			log.Fatalf("creat index.html: %v", err)
		}
		preCount := i * r.Site.Limit
		count := preCount + r.Site.Limit
		if i == indexCount-1 {
			if rest != 0 {
				count = preCount + rest
			}
			next = ""
		}
		if i == 1 {
			prev = "index.html"
		}
		var data = make(map[string]interface{})
		data["Articles"] = articles[preCount:count]
		data["Site"] = r.Site
		data["Prev"] = prev
		data["Next"] = next
		data["Title"] = r.Site.Title
		err = tpl.Execute(outfile, data)
		if err != nil {
			log.Fatalf("Execute %v: %v", link, err)
		}
	}
}

func (r *Render) Archive(tpl template.Template, articles library.Articles) {
	link := filepath.Join(r.Path, "archive.html")
	outfile, err := os.Create(link)
	defer outfile.Close()
	if err != nil {
		log.Fatalf("creat archive.html: %v", err)
	}

	var articleMap = make(map[int][]library.Article)
	for _, article := range articles {
		articleMap[article.Date.Year()] = append(articleMap[article.Date.Year()], article)
	}

	var archives library.Archives
	for year, articlesT := range articleMap {
		archive := library.Archive{
			Year:     year,
			Articles: articlesT,
		}
		archives = append(archives, archive)
	}
	var data = make(map[string]interface{})
	data["Archives"] = archives
	data["Site"] = r.Site
	data["Title"] = "Archive"

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
	data["Title"] = "About Me"

	err = tpl.Execute(outfile, data)
	if err != nil {
		log.Fatalf("Execute about.html: %v", err)
	}
}
