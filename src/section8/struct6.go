package main

import (
	"fmt"
)

type Car struct { // 첫 글자 대문자로 선언(public)
	name    string "차랑명"
	color   string "색상"
	company string "제조사"
	detail spec "상세"
}

type spec struct { // 첫 글자 소문자로 선언(private)
	length int "전장"
	height int "전고"
	width  int "전축"
}

func main() {
	// 중첩 구조체
	car1 := Car{
		"Soul",
		"grey",
		"hyundai",
		spec{100, 100, 200},
	}

	fmt.Println("car1 : ", car1)

	fmt.Println("length : ", car1.detail.length)
	fmt.Println("width : ", car1.detail.width)
}
