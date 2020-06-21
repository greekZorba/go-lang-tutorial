// 대상 사이트 : 루리웹
package main

import (
	_ "bufio"
	"fmt"
	"net/http"
	_ "os"
	"strings"
	_ "strings"
	"sync"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// yhat/scrape : Go simple scrape pkg - 사용하기 어렵지만 코드 학습 위해서 사용
// go-colly : Go scrap and crawling library - goquery 기반 강력하고 쉬운 패키지(많이 사용)
// PuerkitoBio/goquery : 쉬운 html 파싱 지원

// 스크랩핑 대상 url
const (
	urlRoot = "https://ruliweb.com"
)

// 첫 번째 방문(메인 페이지) 대상으로 원하는 url을 파싱 후 반환하는 함수
func parseMainNodes(n *html.Node) bool {
	if n.DataAtom == atom.A && n.Parent != nil {
		return scrape.Attr(n.Parent, "class") == "row"
	}

	return false
}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// 동기화를 위한 작업 그룹 선언
var workGroup sync.WaitGroup

func main() {
	// 메인 페이지 GET 방식 요청
	response, err := http.Get(urlRoot)
	errorCheck(err)

	// 요청 Body 닫기
	defer response.Body.Close()

	// 응답 데이터(Html)
	root, err := html.Parse(response.Body)
	errorCheck(err)

	// ParseMainNodes 메서드를 크롤링(스크랩핑) 대상 URL 추출
	urlList := scrape.FindAll(root, parseMainNodes)

	for _ , link := range urlList {
		// 대상 url 1차 출력
		//fmt.Println("check main link : ", idx, *link)
		//fmt.Println("Target url : ", scrape.Attr(link, "href"))
		fileName := strings.Replace(scrape.Attr(link, "href"), "https://bbs.ruliweb.com/family/", "", 1)
		//fmt.Println("fileName : ", fileName)

		// 작업 대기열에 추가
		workGroup.Add(1)
		
	}
}
