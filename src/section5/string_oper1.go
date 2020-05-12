package main

import "fmt"

func main() {
	// 문자열 연산
	// 추출, 비교, 결합(조합)

	var str1 string = "Golang"
	var str2 string = "World"

	fmt.Println(str1[0:2], str1[0])
	fmt.Println(str2[3:], str2[0])

	fmt.Println(str1[:3])
	fmt.Println(str2[1:3] )
}
