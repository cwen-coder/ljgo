package lib

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	CONFIG_SPLIT = "---"
	MORE_SPLIT   = "<!--more-->"
)

type ConfigArticle struct {
	Title   string
	Date    time.Time
	Update  time.Time
	Tags    []string
	Content string
}

type Article struct {
	SiteConfig
	ConfigArticle
	AuthorConfig
	Preview template.HTML
	Content template.HTML
	Link    string
}

func (a *Article) ParseArticle(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read article: %v", err)
	}
	dataStr := string(data)
	markdownStr := strings.SplitN(dataStr, CONFIG_SPLIT, 2)
	dataLen := len(markdownStr)

	var configStr string
	var contentStr string
	if dataLen > 0 {
		configStr = markdownStr[0]
	}

	if dataLen > 1 {
		contentStr = markdownStr[1]
	}

	if err = yaml.Unmarshal([]byte(configStr), &a.ConfigArticle); err != nil {
		return fmt.Errorf("Unmarshal configArticle: %v", err)
	}

	return nil
}
