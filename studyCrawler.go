/*
대상 웹 사이트 : 정부 코로나 현황
기능 : 1. 뉴스, 보도자료 링크 및 텍스트 파싱 연습
*/
package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "http://ncov.mohw.go.kr/"
	res, _ := http.Get(url)
	doc, _ := goquery.NewDocumentFromResponse(res)

	//	fmt.Println(doc.Find("h2.title1").Text()) // 태그 뒤에 class면 . ID면 #
	//fmt.Println(doc.Find("div.occurrenceStatus").Text())

	//링크 출력 연습, for문이 아니라 Each로 반복문 돌려야 됌
	//	fmt.Println(doc.Find("ul.m_text_list a[href]").Text())
	links := doc.Find("ul.m_text_list a[href]")
	links.Each(func(idx int, sel *goquery.Selection) {
		link, _ := sel.Attr("href")
		text := sel.Text()
		fmt.Println(url+link, text)
	})
}
