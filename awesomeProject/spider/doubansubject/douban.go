package doubansubject

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/lyloou/goer/awesomeProject/spider/util"
)

type Book struct {
	Id    string
	Title string
	Pic   string
	Info  string
}
type Douban struct {
}

func NewDouban() *Douban {
	return &Douban{}
}

func (that *Douban) Finish() {
	fmt.Println("done.")
}

func (that *Douban) Process(p *page.Page) {
	if !p.IsSucc() {
		println(p.Errormsg())
		return
	}

	doc := p.GetHtmlParser()
	tmp, err := url.Parse(p.GetRequest().Url)
	if err != nil {
		panic(err)
	}
	doc.Url = tmp // can not get url by doc.Url directly, need reassign from p.GetRequest().Url
	info := GetDoubanInfo(doc)

	// save info to DB

	fmt.Println(fmt.Sprintf("====>%#v", info))
}

func GetDoubanInfo(doc *goquery.Document) *Book {
	book := &Book{}
	id := getBookId(doc)
	book.Id = id

	wrapper := doc.Find("body > #wrapper")
	title := wrapper.Find("h1")
	titleText := util.GetSelectionText(title)
	book.Title = titleText

	article := wrapper.Find(".article").First()
	mainPic, _ := article.Find("#mainpic .nbg").Attr("href")
	book.Pic = mainPic

	info := article.Find("#info")
	book.Info = util.GetSelectionText(info)

	return book
}

func getBookId(doc *goquery.Document) string {
	split := strings.Split(doc.Url.String(), "/")
	return split[len(split)-2]
}
