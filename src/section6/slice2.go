package main

import "fmt"

func main() {
	// 슬라이스(슬라이스 참조 타입 증명)

	// 배열은 값이 복사됨
	array1 := [3]int{1, 2, 3}
	var array2 [3]int

	array2 = array1
	array2[0] = 20

	fmt.Println("array1 : ", array1)
	fmt.Println("array2 : ", array2)

	slice3 := []int{1, 2, 3}
	var slice4 []int

	slice4 = slice3
	slice4[0] = 20

	fmt.Println("slice3 : ", slice3)
	fmt.Println("slice4 : ", slice4)

	slice5 := make([]int, 50, 100)
	slice5[20] = 30

	fmt.Println("slice5 : ", slice5)
	//fmt.Println("slice5[50] : ", slice5[50]) index out of range

	// 슬라이스 순회
	for index, value := range slice5{
		fmt.Println("index : ", index, "value : ", value)
	}
}
