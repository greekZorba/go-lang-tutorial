package main

import "fmt"

func main() {
	i := 7
	var p = &i

	fmt.Println("i :", i, " *p :", *p, " &p :", &p)

	*p++

	fmt.Println("i : ", i)
	fmt.Println("p : ", *p)

	*p = 18 // 포인터 역참조값 변경

	fmt.Println("i : ", i)
	fmt.Println("p : ", *p)
}
