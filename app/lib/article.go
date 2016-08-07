package lib

import (
	"html/template"
	"time"
)

const (
	CONFIG_SPLIT = "---"
	MORE_SPLIT   = "<!--more-->"
)

type MarkdownArticle struct {
	Title   string
	Date    time.Time
	Update  time.Time
	Preview string
	Tags    []string
	Content string
}

func (m *MarkdownArticle) ParseArticle(path string) (*Article, error) {
	//data, err := ioutil.ReadFile(path)
	//if err != nil {
	//return nil, fmt.Errorf("read article: %v", err)
	//}
}

type Article struct {
	SiteConfig
	MarkdownArticle
	AuthorConfig
	Preview template.HTML
	Content template.HTML
	Link    string
}
