package main

import "fmt"

func main() {
	// 제어문(조건문)
	// if문 : 반드시 Boolean 검사
	// 소괄호를 사용하지 않음

	var a int = 15
	b := 16

	if b > a {
		fmt.Println("b가 a보다 크다")
	}

	if c := 40; c > 30 {
		fmt.Println("true")
	}

	// 중괄호가 if랑 같은 라인에 둬야함, 아래 같은 경우는 에러
	// if a > b
	// {
	// }

	// 중괄호 생략하면 에러
	// if a > b
	// 	fmt.Println("print")

	// boolean 타입이 되어야함
	// if c := 1; c {
	// 	fmt.Println("true")
	// }
}
