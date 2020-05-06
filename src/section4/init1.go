package main

import (
	"fmt"
)

// main 함수보다 먼저 실행됨
func init() {
	fmt.Println("Init1 method start!")
}

func main() {
	// init: 패키지 로드 시에 가장 먼저 호출
	// 가장 먼저 초기화 되는 작업 작성 시 유용하다
	fmt.Println("Main method start!")
}
