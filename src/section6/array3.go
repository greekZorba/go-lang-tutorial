package main

import "fmt"

func main() {
	// 배열 복사

	array1 := [4]int{1, 10, 1000, 10000}
	array2 := array1

	fmt.Println("array1 : ", array1)
	fmt.Println("array2 : ", array2)

	fmt.Printf("array1 pointer : %p value : %v \n", &array1, array1)
	fmt.Printf("array2 pointer : %p value : %v \n", &array2, array2)
}
