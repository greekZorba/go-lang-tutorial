package main

import "fmt"

func ex_f1() {
	fmt.Println("ex_f1 starts!!")
	defer ex_f2() // 마지막에 호출
	fmt.Println("ex_f1 ends!!")
}

func ex_f2() {
	fmt.Println("ex_f2 called")
}

func main() {
	// Defer 함수 실행(지연)
	// Defer를 호출한 함수가 종료되기 직전에 호출된다
	// 타 언어의 finally문과 비슷
	// 주요 리소스 반환, 열린 파일 닫기, Mutex 잠금 해제

	ex_f1()
}
