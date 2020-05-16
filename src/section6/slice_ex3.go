package main

import (
	"fmt"
	"sort"
)

func main() {
	// 슬라이스 추출 및 정렬
	// slice[i:j] i -> j-1 까지 추출
	// slice[i:] i -> 마지막까지 추출
	// slice[:j] 처음 -> j-1까지 추출
	// slice[:] 처음 -> 마지막까지 추출

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("slice[i:j] : ", slice1[1:5])
	fmt.Println("slice[i:] : ", slice1[2:])
	fmt.Println("slice[:j] : ", slice1[:8])
	fmt.Println("slice[:] : ", slice1[:])
	fmt.Println("slice[i:len(slice)]", slice1[0:len(slice1)])

	// sort 패키지
	slice2 := []int{4, 6, 1, 45, 20, 12, 53, 100, 2}

	fmt.Println(sort.IntsAreSorted(slice2)) // 정렬되어있는지 체크
	sort.Ints(slice2) // 정렬
	fmt.Println(slice2)

	slice3 := []string{"z", "j", "a", "f", "h", "x"}
	fmt.Println(sort.StringsAreSorted(slice3))
	sort.Strings(slice3)
	fmt.Println(slice3)
	fmt.Println(sort.StringsAreSorted(slice3))
}
