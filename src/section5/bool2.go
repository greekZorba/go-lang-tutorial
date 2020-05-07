package main

import "fmt"

func main() {
	var b1 bool = true
	var b2 bool = false
	b3 := true

	// 논리 연산자
	fmt.Println("b1 && b2 :", b1 && b2)
	fmt.Println("b1 || b2 :", b1 || b2)
	fmt.Println("!b3 :", !b3)

	num1 := 15
	num2 := 37

	// 비교 연산자
	fmt.Println("num1 > num2 : ", num1 > num2)
	fmt.Println("num1 == num2 : ", num1 == num2)
	fmt.Println("num1 != num2 : ", num1 != num2)
	fmt.Println("num1 < num2 : ", num1 < num2)
}
