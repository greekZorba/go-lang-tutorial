package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

func printValue(s interface{}) {
	fmt.Println("s : ", s)
}

func main() {
	// 인터페이스 활용(빈 인터페이스)
	// 함수 내에서 어떠한 타입이라도 유연하게 매개변수로 받을 수 있다(만능) -> 모든 타입 지정 가능

	dog1 := Dog{"show", 5}
	cat1 := Cat{"myo", 7}

	printValue(dog1)
	printValue(cat1)
	printValue("뭐야")
	printValue(12)
	printValue(1.73)
	printValue(map[string]string{
		"daum":   "daum.com",
		"naver":  "naver.com",
		"google": "google.com",
	})

}
