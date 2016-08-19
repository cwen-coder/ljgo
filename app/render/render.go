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
		link := filepath.Join(pubilcPath, article.ConfigArticle.Title)
		outfile, err := os.Create(link)
		if err != nil {
			log.Fatal(err)
		}

		err = tpl.Execute(outfile, article)
		if err != nil {
			log.Fatal(err)
		}
	}
}
