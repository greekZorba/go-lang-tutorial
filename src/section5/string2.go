package main

import "fmt"

func main() {
	// 문자열 표현
	// 문자열 : utf-8 인코딩 (유니코드 문자 집합) -> 바이트 수 주의!

	var str1 string = "Golang"
	var str2 string = "world"
	var str3 string = "고랭"

	fmt.Println("str1 : ", str1[0], str1[1])
	fmt.Println("str2 : ", str2[0], str2[1])
	fmt.Println("str3 : ", str3[0], str3[1])

	fmt.Printf("str1 : %c %c \n", str1[0], str1[1])
	fmt.Printf("str2 : %c %c \n", str2[0], str2[1])
	fmt.Printf("str3 : %c %c \n", str3[0], str3[1]) // 한글 깨짐

	conStr := []rune(str3)
	fmt.Printf("str3 : %c %c \n", conStr[0], conStr[1]) // 한글 정상 출력

	for i, char := range str1 {
		fmt.Printf("example :  %c(%d)\t", char, i)
	}
}
