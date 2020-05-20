package main

import "fmt"

func multiply(n ...int) int {
	tot := 1
	for _, value := range n {
		tot *= value
	}

	return tot
}

func sum(n ...int) int {
	var result int

	for _, value := range n {
		result += value
	}

	return result
}

func printWord(message ...string) {

	for _, value := range message {
		fmt.Println("message : ", value)
	}

}

func main() {
	// 가변 인자 실습(매개 변수 개수가 동적으로 변할때 - 정해져 있지 않음)
	result := multiply(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("result : ", result)

	sumResult := sum(2, 3, 6, 1, 3, 2, 12, 5)
	fmt.Println("sumResult : ", sumResult)

	printWord("heoo", "wow", "bowwow", "what ? ")

	a := []int{1, 2, 3, 5, 6, 1}
	arrayMultiplyResult := multiply(a...)
	arraySumResult := sum(a...)
	fmt.Println("arrayMultiplyResult : ", arrayMultiplyResult)
	fmt.Println("arraySumResult : ", arraySumResult)
}
