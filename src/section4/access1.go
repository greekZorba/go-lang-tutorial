package main

import (
	"fmt"
	"section4/lib2"
)

func main() {
	// 패키지 접근 제어
	// 변수, 상수, 메서드, 구조체 등 식별자
	// 대문자 : 패키지 외부에서 접근 가능
	// 소문자 : 패키지 외부 접근 불가(패키지 내에서만 접근 가능)
	fmt.Println(lib2.CheckNum1(200))
	fmt.Println(lib2.CheckNum2(1100))
}
