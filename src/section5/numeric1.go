package main

import "fmt"

func main() {
	// 데이터 타입 : 숫자형
	// 정수, 실수, 복소수
	// 32bit, 64bit, unsigned(양수)
	// 정수: 8진수(0), 16진수(0x), 10진수

	var num1 int = 17
	var num2 int = -35
	var num3 int = 0122 // 8진수
	var num4 int = 0x32fa2c75 // 16진수

	fmt.Println("num1 : ", num1)
	fmt.Println("num2 : ", num2)
	fmt.Println("num3 : ", num3)
	fmt.Println("num4 : ", num4)

}
