package main

import (
	oper "./arithmetics" // alias 사용
	"fmt"
)

func main() {
	// 사용자 패키지 작성 & Go 문서화
	// 기준은 GOPATH/src
	// 패키지 폴더명(디렉토리명) 명확하게 지정
	// 패키지 파일의 package 이름으로 사용한다 -> 길면 alias 사용
	// package main을 제외하고 package 문서에 등록
	// 기본적으로 GOROOT의 패키지(pkg) 검색하고 없으면 -> GOPATH의 패키지(src/pkg) 검색
	// go install 명령어 실행의 경우 -> GOPATH/pkg에 등록(문서화)
	// godoc --http=6060(임의의 포트) -> pkg 이동 -> 본인 패키지 메서드 및 주석 확인(패키지, 타입, 메서드) 주석

	// 패키지 사용 예제(사칙연산)
	nums := oper.Numbers{X: 10, Y: 10}
	fmt.Println("Package Used(1) : ", nums.Plus())
	fmt.Println("Package Used(2) : ", nums.Minus())
	fmt.Println("Package Used(3) : ", nums.Multi())
	fmt.Println("Package Used(4) : ", nums.Divide())
	fmt.Println("Package Used(5) : ", nums.SquarePlus())
	fmt.Println("Package Used(6) : ", nums.SquareMinus())
}
