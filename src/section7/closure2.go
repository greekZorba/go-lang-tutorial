package main

import "fmt"

func main() {
	cnt := increaseCnt()

	fmt.Println("cnt : ", cnt())
	fmt.Println("cnt : ", cnt())
	fmt.Println("cnt : ", cnt())
	fmt.Println("cnt : ", cnt())
	fmt.Println("cnt : ", cnt())
	fmt.Println("cnt : ", cnt())
	fmt.Println("cnt : ", cnt())
}

func increaseCnt() func() int {
	n := 0 // 지역변수(캡처됨)
	return func() int {
		n += 1
		return n
	}
}
