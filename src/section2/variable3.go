package main

import "fmt"

func main() {
	// 짧은 선언
	// 함수 안에서만 사용(전역x), 선언 후 할당 예외 발생
	// 주로 제한된 범위내에서 사용할 경우 코드 가독성을 높일 수 있다 -> 추천
	shortVar1 := 1
	shortVar2 := "Test"
	shortVar3 := false

	// := 선언 후 재할당시 예외 발생, 컴파일 에러
	fmt.Println("shortVar1 : ", shortVar1, "shortVar2 : ", shortVar2, "shortVar3 : ", shortVar3)

	if i := 10; i < 11 {
		fmt.Println("10 smaller than 11 ")
	}
}
