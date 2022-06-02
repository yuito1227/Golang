package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://www.jma.go.jp/bosai/forecast/#area_type=offices&area_code=270000")
	if err != nil {
		panic("htmlの取得に失敗しました")
	}
	title := doc.Find(".forecast-sentence")
	fmt.Print(title)
	title.Each(func(i int, s *goquery.Selection) {
		fmt.Print(s.Text())
	})
}
