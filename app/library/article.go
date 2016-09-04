package library

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"git.cwengo.com/cwen/ljgo/app/util"

	"gopkg.in/yaml.v2"
)

const (
	CONFIG_SPLIT = "---"
	MORE_SPLIT   = "<!--more-->"
)

type ConfigArticle struct {
	Title  string   `yaml:"title"`
	Date   string   `yaml:"date"`
	Update string   `yaml:"update"`
	Tags   []string `yaml:"tags"`
	Author string   `yaml:"author"`
}

type Article struct {
	ConfigArticle
	// AuthorConfig
	Date    time.Time
	Update  time.Time
	Preview template.HTML
	Content template.HTML
	Link    string
}

func NewArticle() *Article {
	return &Article{}
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
	err = a.ParseDate(a.ConfigArticle.Date, a.ConfigArticle.Update)
	if err != nil {
		return err
	}
	year := strconv.Itoa(a.Date.Year())
	a.Link = year + "/" + a.ConfigArticle.Title + ".html"

	a.ParseMarkdown(contentStr)

	return nil
}

func (a *Article) ParseDate(date, update string) error {

	var err error
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}
	a.Date, err = util.ParseDate(date)
	if err != nil {
		return fmt.Errorf("parse date: %v", err)
	}
	if update == "" {
		update = date
	}
	a.Update, err = util.ParseDate(update)
	if err != nil {
		return fmt.Errorf("parse update: %v", err)
	}
	return nil
}

func (a *Article) ParseMarkdown(contentStr string) {
	contentArr := strings.SplitN(contentStr, MORE_SPLIT, 2)
	if len(contentArr) > 1 {
		a.Preview = util.ParseMarkdown(contentArr[0])
	}

	contentStr = strings.Replace(contentStr, MORE_SPLIT, "", 1)

	a.Content = util.ParseMarkdown(contentStr)
}

type Articles []Article

func (a Articles) Len() int {
	return len(a)
}

func (a Articles) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Articles) Less(i, j int) bool {
	return a[i].Date.Unix() > a[j].Date.Unix()
}
