package render

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/feeds"
	"github.com/qiniu/log"

	"git.cwengo.com/cwen/ljgo/app/config"
	"git.cwengo.com/cwen/ljgo/app/library"
	"git.cwengo.com/cwen/ljgo/app/util"
)

type Render struct {
	Site config.SiteConfig
	Path string
}

func New(cfg *config.Config) *Render {
	return &Render{
		Site: cfg.Site,
		Path: cfg.PublicPath,
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
		err := os.MkdirAll(r.Path+"/"+strconv.Itoa(year), 0777)
		if err != nil {
			log.Fatalf("mkdir %v: %v", year, err)
		}
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

func (r *Render) Tags(tpl template.Template, articles library.Articles) {
	var TagsMap = make(map[string]library.Articles)
	for _, article := range articles {
		for _, tag := range article.ConfigArticle.Tags {
			TagsMap[tag] = append(TagsMap[tag], article)
		}
	}
	err := os.MkdirAll(filepath.Join(r.Path, "tags"), 0777)
	if err != nil {
		log.Fatalf("madir tag: %v", err)
	}
	for tag, articlesT := range TagsMap {
		link := filepath.Join(r.Path, "tags/"+tag+".html")
		outfile, err := os.Create(link)
		defer outfile.Close()
		if err != nil {
			log.Fatalf("creat article %v: %v", link, err)
		}
		var data = make(map[string]interface{})
		data["Tag"] = tag
		data["Articles"] = articlesT
		data["Site"] = r.Site
		err = tpl.Execute(outfile, data)
		if err != nil {
			log.Fatalf("Execute %v.html: %v", tag, err)
		}
	}
}

func (r *Render) RSS(articles library.Articles) {
	var feedArticles library.Articles
	if len(articles) > r.Site.Limit {
		feedArticles = articles[0:r.Site.Limit]
	} else {
		feedArticles = articles
	}
	if r.Site.URL == "" {
		return
	}
	feed := &feeds.Feed{
		Title:       r.Site.Title,
		Link:        &feeds.Link{Href: r.Site.URL},
		Description: r.Site.Introduce,
		Created:     time.Now(),
		Items:       make([]*feeds.Item, 0),
	}
	for _, item := range feedArticles {
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       item.ConfigArticle.Title,
			Link:        &feeds.Link{Href: r.Site.URL + "/" + item.Link},
			Description: string(item.Content),
			Author:      &feeds.Author{Name: item.ConfigArticle.Author},
			Created:     item.Date,
			Updated:     item.Update,
		})
	}
	atom, err := feed.ToAtom()
	if err != nil {
		log.Fatalf("rss ToAtom: %v", err)
	}
	path := filepath.Join(r.Path, "atom.xml")
	if _, err := os.Stat(path); err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("open atom.xml: %v", err)
		}
		file, err := os.Create(path)
		if err != nil {
			log.Fatalf("create atom.xml: %v", err)
		}
		file.Close()
	}
	err = ioutil.WriteFile(path, []byte(atom), 0644)
	if err != nil {
		log.Fatalf("write atom.xml: %v", err)
	}
}
