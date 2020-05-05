package main

import (
	"fmt"
	checkUp10 "section4/lib"
	_ "section4/lib2"
)

func main() {
	// 패키지 접근 제어
	// 별칭 사용 ( "_"을 사용해서 사용하지 않아도 그대로 둘 수 있음)
	// 빈 식별자
	fmt.Println(checkUp10.CheckNum(100))
}
