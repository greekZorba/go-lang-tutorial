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

	
}
