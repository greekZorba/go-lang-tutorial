package main

import "fmt"

func main() {
	a := 30 / 15
	switch a {
	case 2, 4, 6:
		fmt.Println("a -> ", a, "는 짝수다")
	case 1, 3, 5:
		fmt.Println("a -> ", a, "는 홀수다")
	}

	// fallthrough 조건문 무시하고 다음 조건문을 실행시킴
	switch e := "Go"; e {
	case "Java":
		fmt.Println("This is Java!")
	case "Go":
		fmt.Println("This is Go!")
		fallthrough
	case "Python":
		fmt.Println("This is Python!")
		fallthrough
	case "C":
		fmt.Println("This is C!")
	}
}
