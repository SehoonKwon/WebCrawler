/*
대상 웹 사이트 : 정부 코로나 현황
기능:	1. 뉴스, 보도자료 링크 및 텍스트 파싱 연습
		2. 예방접종현황 파싱 (ing)
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

	/*
		//링크 출력 연습, for문이 아니라 Each로 반복문 돌려야 됌
		//	fmt.Println(doc.Find("ul.m_text_list a[href]").Text())
		links := doc.Find("ul.m_text_list a[href]")
		links.Each(func(idx int, sel *goquery.Selection) {
			link, _ := sel.Attr("href")
			text := sel.Text()
			fmt.Println(url+link, text)
		})
	*/

	//예방접종현황
	var vaccination []string
	var percentList []string
	var personCnt []string

	vaccine_list := doc.Find("div.vaccine_list")
	vaccine_list.Each(func(idx int, sel *goquery.Selection) {
		/*
			fmt.Println(sel.Find("div.item").Text())
			fmt.Println(sel.Find("li.percent").Text())
			fmt.Println(sel.Find("li.person").Text())
		*/
		vac := sel.Find("div.item")
		perc := sel.Find("li.percent")
		per := sel.Find("li.person")

		vac.Each(func(idx int, thisSel *goquery.Selection) {
			vaccination = append(vaccination, thisSel.Text())
		})

		perc.Each(func(idx int, thisSel *goquery.Selection) {
			percentList = append(percentList, thisSel.Text())
		})

		per.Each(func(idx int, thisSel *goquery.Selection) {
			personCnt = append(personCnt, thisSel.Text())
		})
	})

	/*
		var vaccination []string
		parse_vaccination := doc.Find("div.vaccine_list div.item")
		parse_vaccination.Each(func(idx int, sel *goquery.Selection) {
			vaccination = append(vaccination, sel.Text())
		})
	*/

	fmt.Println(vaccination)
	fmt.Println(percentList)
	fmt.Println(personCnt)
}
