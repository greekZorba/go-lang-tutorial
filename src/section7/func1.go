package main

import (
	"fmt"
	"strconv"
)

func helloGolang() {
	fmt.Println("Hello! Golang!")
}

func say_one(message string) {
	fmt.Println("received message : ", message)
}

func sum(x int, y int) int {
	return x + y
}

func main() {
	// func 함수명(매개변수) 반환타입 or 반환 값 변수명 : 매개변수 있고, 반환값 존재
	// func 함수명() 반환타입 or 반환 값 변수명 : 매개변수 없고, 반환값 존재
	// func 함수명(매개변수) : 매개변수 존재, 반환값 없음
	// func 함수명() : 매개변수 없음, 반환값 없음
	// 타 언어와 달리 반환값(return value) 다수 가능

	helloGolang()
	say_one("Hellowwww~!!")
	fmt.Println("sum : ", sum(10, 20))

	// 숫자를 문자열로 변환해줌
	fmt.Println("int to string ", strconv.Itoa(sum(20, 20)))
}
