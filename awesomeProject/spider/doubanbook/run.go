package doubanbook

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
)

func RunSingle() {
	id := "4820711"
	nd := NewDouban()
	spider.NewSpider(nd, "douban").
		AddUrl("https://book.douban.com/subject/"+id+"/", "html").
		AddPipeline(pipeline.NewPipelineConsole()).
		SetThreadnum(3).
		Run()
}

func RunMultiple() {
	start := 4820711
	nd := NewDouban()
	urls := make([]string, 10)
	for i := 0; i < 10; i++ {
		id := strconv.Itoa(start + i)
		url := fmt.Sprintf("https://book.douban.com/subject/%s/", id)
		urls = append(urls, url)
	}
	spider.NewSpider(nd, "douban").
		AddUrls(urls, "html").
		AddPipeline(pipeline.NewPipelineConsole()).
		SetThreadnum(3).
		Run()
}

func RunMultipleWithDefer() {
	start := 4820711
	ticker := time.NewTicker(time.Second * 3)
	nd := NewDouban()

	for range ticker.C {
		start++
		url := fmt.Sprintf("https://book.douban.com/subject/%s/", strconv.Itoa(start))
		spider.NewSpider(nd, "douban").
			AddUrl(url, "html").
			AddPipeline(pipeline.NewPipelineConsole()).
			SetThreadnum(1).
			Run()
	}
}
