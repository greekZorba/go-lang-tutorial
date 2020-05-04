package main

import "fmt"

func main() {
	// 코드 서식 지정
	// 한 사람이 코딩한 것 같은 일관성 유지
	// 코드 스타일 유지
	// gofmt -h : 사용법
	// gofmt -w : 원본파일에 반영

	for i := 0; i < 10; i++ {
		fmt.Println("i :", i)
	}
}
