package util

import (
	"strings"
	"regexp"
	"github.com/grokify/html-strip-tags-go"
	"github.com/PuerkitoBio/goquery"
)

func ReplaceHtml(string string) string {
	string = strings.Replace(string, "\n", "", -1)
	string = strings.Replace(string, "\t", "", -1)
	return string
}

func ReplaceString(s string, old string) string {
	s = strings.Replace(s, old, "", -1)
	return s
}

func RemoveDuplicateWhitespace(src string) string {

	space := regexp.MustCompile(`\s+`)
	dst := space.ReplaceAllString(src, " ")
	dst = strings.TrimSpace(dst)
	return dst
}

func GetSelectionText(selection *goquery.Selection) string {
	tmp, _ := selection.Html()
	tmp = ReplaceHtml(tmp)
	tmp = strings.Replace(tmp, "：", "", -1)
	tmp = strip.StripTags(tmp)
	return RemoveDuplicateWhitespace(tmp)
}
