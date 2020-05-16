package main

import "fmt"

func main() {
	// slice 추가 및 병합

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}
	slice3 := []int{11, 12, 13}

	slice1 = append(slice1, 9, 10, 11)
	slice2 = append(slice1, slice2...) // 슬라이스를 삽입할 경우 ... 사용
	slice3 = append(slice2, slice3[1:3]...) // 추출 후 병합

	fmt.Println("modified slice1 : ", slice1)
	fmt.Println("modified slice2 : ", slice2)
	fmt.Println("modified slice3 : ", slice3)

	slice4 := make([]int, 0, 5)

	for i := 0; i < 15; i++{
		slice4 = append(slice4, i)
		// 길이 및 용량 자동증가
		fmt.Printf(" length : %d, capacity : %d, value : %v \n", len(slice4), cap(slice4), slice4[i])
	}


}
