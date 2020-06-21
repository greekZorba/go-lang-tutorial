// 대상 사이트 : 루리웹
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
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

// url 대상이 되는 페이지(서브페이지) 대상으로 원하는 내용을 파싱 후 반환
func scrapContents(url string, fileName string) {
	defer workGroup.Done()

	// get 방식 요청
	response, err := http.Get(url)
	errorCheck(err)

	defer response.Body.Close()

	// 응답 데이터(Html)
	root, err := html.Parse(response.Body)
	errorCheck(err)

	// response 데이터(html)의 원하는 부분 파싱
	matchNode := func(n *html.Node) bool {
		return n.DataAtom == atom.A && scrape.Attr(n, "class") == "deco"
	}

	// 파일 스크림 생성(열기) -> 파일명, 옵션, 권한
	file, err := os.OpenFile("/Users/jinhakkim/Desktop/goCrawler/"+fileName+".txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	errorCheck(err)

	// 메서드 종료시 파일 닫기
	defer file.Close()

	// 쓰기 버퍼 선언
	writer := bufio.NewWriter(file)

	// matchNode를 이용해 원하는 노드 순회(iterator)하면서 출력
	for _, value := range scrape.FindAll(root, matchNode) {
		// url 및 해당 데이터 출력
		fmt.Println("result : ", scrape.Text(value))
		// 파싱 데이터 -> 버퍼에 기록
		writer.WriteString(scrape.Text(value) + "\r\n")
	}

	writer.Flush()
}

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

	for _, link := range urlList {
		// 대상 url 1차 출력
		//fmt.Println("check main link : ", idx, *link)
		//fmt.Println("Target url : ", scrape.Attr(link, "href"))
		fileName := strings.Replace(scrape.Attr(link, "href"), "https://bbs.ruliweb.com/family/", "", 1)
		//fmt.Println("fileName : ", fileName)

		// 작업 대기열에 추가
		workGroup.Add(1) // done 개수와 일치
		// 고루틴 시작 -> 작업 대기열 개수와 같아야함
		go scrapContents(scrape.Attr(link, "href"), fileName)

	}
	// 모든 작업이 종료될 때까지 대기
	workGroup.Wait()
}
