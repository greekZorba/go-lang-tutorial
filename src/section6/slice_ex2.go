package main

import "fmt"

func main() {
	// 슬라이스 복사
	// copy(복사대상, 원본)
	// make로 공간 할당 후 복사해야한다
	// 복사된 슬라이스 값 변경해도 원본에는 영향없음

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{} // 복사 안됨

	copy(slice2, slice1)
	copy(slice3, slice1)

	fmt.Println("slice2 : ", slice2)
	fmt.Println("slice3 : ", slice3)

	slice4 := []int{1, 2, 3, 4, 5}
	slice5 := make([]int, 5)

	copy(slice5, slice4)

	slice5[0] = 100
	slice5[1] = 1000

	fmt.Println("slice4 : ", slice4)
	fmt.Println("slice5 : ", slice5)

	slice6 := [5]int{1, 2, 3, 4, 5}
	slice7 := slice6[0:3] // 주의 : 부분적으로 슬라이스 추출은 참조 -> 원본값 변경됨

	slice7[0] = 100

	fmt.Println("slice6 : ", slice6)
	fmt.Println("slice7 : ", slice7)

	slice8 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice9 := slice8[0:5:7] // 용량 지정

	fmt.Println("slice8 : ", slice8)
	fmt.Println("slice9 : ", slice9, len(slice9), cap(slice9))

}
