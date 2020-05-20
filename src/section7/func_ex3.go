package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func printHello(n int) {
	if n == 0 {
		return
	}

	fmt.Println("n : ", n)
	printHello(n - 1)
}

func main() {
	// 재귀 함수(Recursion)
	// 프로그램이 보기 쉽고, 코드 간결, 오류 수정 용이 : 장점
	// 코드가 이해하기 어렵고, 기억 공간을 많이 사용, 무한루프 가능성

	x := fact(5)
	fmt.Println("factorial : ", x)

	printHello(10)
}
