package main

import "fmt"

func main() {
	// 익명함수 (Anonymous function)
	// 즉시 실행(선언과 동시에)

	func() {
		fmt.Println("Anonymous function")
	}()

	message := "hello world!"

	func(input string) {
		fmt.Println("message : ", input)
	}(message)

	var sumResult = func(x, y int) int {
		return x + y
	}(10, 10)

	fmt.Println("sumResult : ", sumResult)

}
