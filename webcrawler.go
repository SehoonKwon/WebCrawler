//대상 웹 사이트 : 루리웹
//기능 : 각 게임별 게시판 제목 txt 저장
package main

import (
	_ "bufio"
	_ "fmt"
	"net/http"
	_ "os"
	_ "strings"
	"sync"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

//스크래핑 대상 URL
const urlRoot = "http://ruliweb.com"

//	"github.com/yhat/scrape" 예제 참고
// a 태그이면서 부모가 nil이 아닌경우 부모 클래스가 row인 것을 scrape
func parseMainNodes(n *html.Node) bool {
	if n.DataAtom == atom.A && n.Parent != nil {
		return scrape.Attr(n.Parent, "class") == "row"
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
var wg sync.WaitGroup

func main() {

	//메인 페이지 GET 요청
	rep, err := http.Get(urlRoot)
	errcheck(err)

	//요청 Body 닫기
	defer rep.Body.Close()

	//응답 데이터(HTML)
	root, err := html.Parse(rep.Body)
	errcheck(err)

	//Parse Main Nodes 메소드 스크래핑 대상 URL 추출(게임별 세부 URL)
	urlList := scrape.FindAll(rot, pase)
}
