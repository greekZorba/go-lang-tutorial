package main

import "fmt"

func sum(i int, f func(int, int)) {
	f(i, 10)
}

func add(x, y int) {
	fmt.Println("x + y : ", x+y)
}

func multi_value(i int) {
	i = i * 10
}

func multi_reference(i *int) {
	*i = *i * 10
}

func main() {
	// 함수(콜백) 전달, 참조 전달(call by reference), 값 전달(call by value)
	sum(30, add) // 함수 전달

	// call by value
	a := 100
	multi_value(a)
	fmt.Println("a :", a)

	// call by reference
	b := 1000
	multi_reference(&b)
	fmt.Println("b : ", b)

}
