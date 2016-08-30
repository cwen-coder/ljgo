package render

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/qiniu/log"

	"git.cwengo.com/cwen/ljgo/app/library"
)

func RenderArticles(tpl template.Template, articles []library.Article, pubilcPath string) {
	for _, article := range articles {
		link := filepath.Join(pubilcPath, article.Link)
		outfile, err := os.Create(link)
		if err != nil {
			log.Fatalf("creat article %v: %v", link, err)
		}

		err = tpl.Execute(outfile, article)
		if err != nil {
			log.Fatalf("Execute %v: %v", link, err)
		}
	}
}

func RenderIndex(tpl template.Template, articles []library.Article, pubilcPath string) {
	link := filepath.Join(pubilcPath, "index.html")
	outfile, err := os.Create(link)
	if err != nil {
		log.Fatalf("creat index.html: %v", err)
	}
	var data = make(map[string]interface{})
	data["Articles"] = articles

	err = tpl.Execute(outfile, data)
	if err != nil {
		log.Fatalf("Execute index.html: %v", err)
	}
}
