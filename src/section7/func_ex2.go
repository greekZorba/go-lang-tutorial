package main

import "fmt"

func multiply(x, y int) (result int) {
	result = x * y
	return result
}

func sum(x, y int) (result int) {
	result = x + y
	return result
}

func main() {
	// 함수를 변수에 할당

	// 슬라이스에 할당
	f := []func(int, int) int{multiply, sum}

	multiplyResult := f[0](12, 12)
	sumResult := f[1](15, 15)

	fmt.Println("multiplyResult : ", multiplyResult)
	fmt.Println("sumResult : ", sumResult)

	var f1 func(int, int) int = multiply
	f2 := sum

	fmt.Println("f1 : ", f1(10, 20))
	fmt.Println("f2 : ", f2(100, 122))

	m := map[string]func(int, int) int{
		"multiply": multiply,
		"sum":      sum,
	}

	fmt.Println("multiply in map : ", m["multiply"](10, 20))
}
