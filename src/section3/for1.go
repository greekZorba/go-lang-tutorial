package main

import "fmt"

func main() {
	// for문 : Go에서 유일하게 제공하는 반복문
	// 다양한 사용법 숙지

	for i := 0; i < 5; i++ {
		fmt.Println("i : ", i)
	}

	// 에러 발생 1
	/*
		i++ 뒤에 ';'(세미콜론)를 사용해서 안되고, '{'를 줄바꿈해서 안됨
		for i := 0; i < 5; i++;
		{
			fmt.Println("i : ", i)
		}
	*/

	// 에러 발생 2
	/*
		중괄호가 없어서 에러
		for i := 0; i < 5; i++;
			fmt.Println("i : ", i)
	*/

	/*
		무한 루프
		for {

		}
	*/

	i := 0
	for i < 10 {
		i++
		fmt.Println("i : ", i, "무한 루프")
	}

	// Range 용법
	location := []string{"Seoul", "Busan", "Ulsan", "Gangwon"}
	for index, name := range location {
		fmt.Println(index, "번째 도시는 ", name)
	}

	for _, name := range location {
		fmt.Println("도시는 ", name)
	}
}
