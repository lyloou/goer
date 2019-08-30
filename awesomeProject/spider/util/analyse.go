package util

import (
	"container/list"
	"fmt"
	"getweb/core"

	"github.com/PuerkitoBio/goquery"
)

func GetVertical(datalist *goquery.Selection) *list.List {
	if nil == datalist {
		return nil
	}

	th := datalist.Find("tbody tr th")
	thMap := make(map[string]int)

	th.Each(func(i int, t *goquery.Selection) {
		thMap[t.Text()] = i
	})

	l := list.New()
	tr := datalist.Find("tbody tr")
	tr.Each(func(_ int, tr0 *goquery.Selection) {
		tdMap := make(map[string]interface{})
		assigned := false
		tr0.Find("td").Each(func(i int, td *goquery.Selection) {
			key := getIndexByValue(thMap, i)
			value := core.GetSelectionText(td)
			tdMap[key] = value
			assigned = true
		})
		if assigned {
			l.PushBack(tdMap)
		}
	})

	return l
}
func getIndexByValue(m map[string]int, value int) string {
	for k, v := range m {
		if v == value {
			return k
		}
	}
	return ""
}

func GetHorizontal(datalist *goquery.Selection) map[string]interface{} {
	// remove the left <table> tag
	datalist.Find("tbody>tr>td>table").Each(func(_ int, selection *goquery.Selection) {
		img, _ := selection.Find("tbody>tr>td>img").First().Attr("src")
		fmt.Println("src:" + img)
		selection.Remove()
	})

	tdMap := make(map[string]interface{})
	tr := datalist.Find("tbody>tr")
	tr.Each(func(_ int, tr0 *goquery.Selection) {
		td0 := tr0.Find("td")
		td1 := td0.Next()
		key := core.GetSelectionText(td0)
		value := core.GetSelectionText(td1)
		tdMap[key] = value
	})
	return tdMap
}
