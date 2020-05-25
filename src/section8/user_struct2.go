package main

import "fmt"

type cnt int

func main() {
	// 기본 자료형 사용자 정의 타입
	a := 15
	fmt.Println("a : ", a)

	var b cnt = 15
	fmt.Println("b : ", b)

	// testConvertT(b) 컴파일 에러남 -> 타입이 달라서 발생
	// 사용자 정의 타입 <-> 기본 타입 : 매개 변수 전달 시에 변환 해야 사용 가능(int(변수))
	testConvertT(int(b))
	testConvertD(b)
}

func testConvertT(a int)  {
	fmt.Println("Default Type : " , a)
}

func testConvertD(a cnt)  {
	fmt.Println("Custom Type : " , a)
}
