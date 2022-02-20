//대상 웹 사이트 : 루리웹
//기능 : 각 게임별 게시판 제목 txt 저장
package main

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

//스크래핑 대상 URL
const urlRoot = "http://ruliweb.com"

//	"github.com/yhat/scrape" 예제 참고
func parseMainNodes(n *html.Node) bool {
	if n.DataAtom == atom.A && n.Parent != nil { //a 태그이면서 부모가 nil이 아닌경우
		return scrape.Attr(n.Parent, "class") == "row" //  부모 클래스가 row인 것을 scrape
	}
	return false
}

//에러체크 공통함수
func errcheck(err error) {
	if err != nil {
		panic(err)
	}
}

//동기화를 위한 작업 그룹 선언
//모든 고루틴이 종료될 때까지 대기할 때 사용한다.
var wg sync.WaitGroup

//URL 대상이 되는 페이지(서브페이지) 대상으로 원하는 내용을 파싱 후 반환
func scrapContents(url string, fn string) {
	//고루틴 작업 종료 알림
	defer wg.Done()

	//Get 요청
	resp, err := http.Get(url)
	errcheck(err)

	//요청 body 닫기
	defer resp.Body.Close()

	//응답 데이터(HTML)
	root, err := html.Parse(resp.Body)
	errcheck(err)

	//response 데이터 원하는 부분 파싱
	//a 태그이면서 class 가 deco인 것들 matcher로 작성
	matchNode := func(n *html.Node) bool {
		return n.DataAtom == atom.A && scrape.Attr(n, "class") == "deco"
	}

	//파일 스크림 생성(열기) -> 파일명, 옵션, 권한
	//권한 : 읽기 4, 쓰기 2, 실행 1. 4+2+1 = 7 이므로 읽기, 쓰기, 실행 모두 사용가능한 권한이 7. 앞의 순서부터 소유자, 그룹 사용자, 기타 사용자 순의 권한이다
	scrapFolder := "C:\\Users\\Sehoon\\go\\src\\github.com\\SehoonKwon\\WebCrawler\\scrape\\"
	file, err := os.OpenFile(scrapFolder+fn+".txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	errcheck(err)

	//메소드 종료시 파일 닫기
	defer file.Close()

	//쓰기 버퍼 선언. 파일에 작성할 버퍼
	w := bufio.NewWriter(file)

	//matchNode 메소드를 사용해 원하는 노드 순회하면서 출력
	for _, g := range scrape.FindAll(root, matchNode) {
		//url 및 해당 데이터 출력
		//	fmt.Println(scrape.Text(g))

		//파싱 데이터 버퍼에 기록
		w.WriteString(scrape.Text(g) + "\r\n")
	}

	w.Flush()
}

func main() {

	//메인 페이지 GET 요청
	resp, err := http.Get(urlRoot)
	errcheck(err)

	//요청 Body 닫기
	defer resp.Body.Close()

	//응답 데이터(HTML)
	//html.Parse : io.reader로부터 HTML에 대한 구문 분석
	root, err := html.Parse(resp.Body)
	errcheck(err)

	//Parse Main Nodes 메소드 스크래핑 대상 URL 추출(게임별 세부 URL)
	//scrape.FindAll : 주어진 html.Node를 탐색하며 Matcher와 일치하는 모든 노드를 반환
	urlList := scrape.FindAll(root, parseMainNodes)

	for _, link := range urlList {
		//대상 url 1차 출력
		//fmt.Println(link, idx)

		//fmt.Println("Target :", scrape.Attr(link, "href"))
		fileName := strings.Replace(scrape.Attr(link, "href"), "https://bbs.ruliweb.com/family/", "", 1) //scrape 한 문자열의 해당부분을 ""로 치환해라. 1번만
		//fmt.Println(fileName)

		//고루틴 작업 대기열에 추가
		wg.Add(1) //Done 개수와 일치

		//고루틴 시작 -> 작업 대기열 개수와 같아야 함, urlList만큼 추가하고 scrapContents 함수에서 wg.Done을 urlList 만큼 한번씩 수행하니 갯수 일치
		go scrapContents(scrape.Attr(link, "href"), fileName)
	}

	//모든 고루틴 작업이 끝날때까지 대기
	wg.Wait()
}
