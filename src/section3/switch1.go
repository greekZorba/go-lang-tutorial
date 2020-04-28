package main

import "fmt"

func main() {
	/* 제어문(조건문) - switch
	switch 뒤 표현식(expression) 생략 가능
	case 뒤 표현식(expression) 사용 가능
	자동 break 때문에 fallThrough가 존재
	Type 분기 -> 값이 아닌 변수 Type으로 분기 가능
	*/

	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "는 0보다 작다")
	case a == 0:
		fmt.Println(a, "는 0과 같다")
	case a > 0:
		fmt.Println(a, "는 0보다 크다")
	}

	switch b := 27; {
	case b < 0:
		fmt.Println(b, "는 0보다 작다")
	case b == 0:
		fmt.Println(b, "는 0과 같다")
	case b > 0:
		fmt.Println(b, "는 0보다 크다")
	}

	switch c := "Go"; c {
	case "Go":
		fmt.Println("This is Go!")
	case "Java":
		fmt.Println("This is Java!")
	default:
		fmt.Println("일치하는 값 없음")
	}

	switch c := "Go"; c + "Lang"{
	case "GoLang":
		fmt.Println("This is GoLang!")
	case "JavaLang":
		fmt.Println("This is JavaLang!")
	default:
		fmt.Println("일치하는 값 없음")
	}

	switch i, j := 10, 20; {
	case i < j:
		fmt.Println("i는 j보다 작다")
	case i < 15:
		fmt.Println("i는 15보다 작다")
	case i > j:
		fmt.Println("i가 j보다 크다")
	default:
		fmt.Println("일치하는 값 없음")
	}


}
