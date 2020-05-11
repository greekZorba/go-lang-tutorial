package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 문자열
	// 큰 따옴표 "", 백스쿼트 ``
	// Golang: 문자 char 타입 존재하지 않음 -> rune(int32)로 문자 코드 값으로 표현
	// 문자 : ''로 작성
	// 자주 사용하는 escape: \\, \', \", \a(콘솔벨), \b(백스페이스), \f(쪽바꿈), \n(줄바꿈), \r(복귀), \t(탭)..

	var str1 string = "c:\\go_study\\src\\" // -> c:\go_study\src\
	str2 := `c:\go_study\src\`

	fmt.Println("str1 : ", str1)
	fmt.Println("str2 : ", str2)

	var str3 string = "Hello world!"
	var str4 string = "안녕하세요"
	var str5 string = "\ud55c\uae00"

	fmt.Println("str3 : ", str3)
	fmt.Println("str4 : ", str4)
	fmt.Println("str5 : ", str5)

	// 길이(바이트 수)
	fmt.Println("str3 len : ", len(str3))
	fmt.Println("str4 len : ", len(str4))
	fmt.Println("str5 len : ", len(str5))

	// 실제 길이
	fmt.Println("str3 문자 길이 : ", utf8.RuneCountInString(str3))
	fmt.Println("str4 문자 길이 : ", utf8.RuneCountInString(str4))

	// len으로 실제 길이 구하기
	fmt.Println("str5 문자 길이 : ", len([]rune(str5)))
}
