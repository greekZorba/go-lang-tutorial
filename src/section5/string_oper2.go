package main

import "fmt"

func main() {
	// 문자열 비교

	var str1 string = "Golang"
	var str2 string = "World"

	fmt.Println("str1 == str2 : ", str1 == str2)
	fmt.Println("str1 != str2 : ", str1 != str2)
	fmt.Println("str1 > str2 : ", str1 > str2)
	fmt.Println("str1 < str2 : ", str1 < str2) // Go 문자열 -> 아스키 코드에 대한 사전식 비교

}
