package util

import (
	"html/template"

	"github.com/russross/blackfriday"
)

func ParseMarkdown(markdown string) template.HTML {
	return template.HTML(blackfriday.MarkdownCommon([]byte(markdown)))
}
