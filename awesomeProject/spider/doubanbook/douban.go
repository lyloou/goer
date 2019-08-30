package doubanbook

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/lyloou/goer/awesomeProject/spider/util"
)

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

	fmt.Println("info===>", info)
}
func GetDoubanInfo(doc *goquery.Document) interface{} {
	id := getId(doc)
	fmt.Println(id)

	wrapper := doc.Find("body > #wrapper")
	title := wrapper.Find("h1")
	titleText := util.GetSelectionText(title)
	fmt.Println(titleText)

	article := wrapper.Find("#article").First()
	mainPic, _ := article.Find("#mainpic .nbgnbg").Attr("href")
	fmt.Println(mainPic)

	return nil
}

func getId(doc *goquery.Document) string {
	split := strings.Split(doc.Url.String(), "/")
	return split[len(split)-2]
}
