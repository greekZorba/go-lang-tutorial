package main

import "fmt"

func main() {
	// 배열 순회

	array1 := [5]int{1, 10, 100, 1000, 10000}

	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}
	fmt.Println()

	for index, value := range array1 {
		fmt.Printf("index : %d value : %d \n", index, value)
	}
	fmt.Println()

	for _, value := range array1 {
		fmt.Printf("value : %d \n", value)
	}
	fmt.Println()

	for index := range array1 {
		fmt.Printf("index : %d \n", index)
	}

}
