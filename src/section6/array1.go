package main

import "fmt"

func main() {
	// 배열
	// 배열은 용량, 길이 항상 같다
	// 배열 vs 슬라이스 차이점 중요
	// 길이고정 vs 길이가변
	// 값 타입 vs 참조 타입
	// 복사 전달 vs 참조 값 전달
	// 전체 비교 연산자 사용 가능 vs 비교 연산자 사용불가
	// 대부분 슬라이스 사용한다

	// cap() : 배열, 슬라이스 용량
	// len() : 배열, 슬라이스 개수

	var array1 [5]int
	var array2 [5]int = [5]int{1, 2, 3, 4, 5}
	var array3 = [5]int{1, 2, 3, 4, 5}
	array4 := [5]int{1, 2, 3, 4, 5}
	array5 := [5]int{1, 2, 3}
	array6 := [...]int{1, 2, 3} // 배열 크기 자동 맞춤
	array7 := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	array1[2] = 5

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)
	fmt.Println(array5)
	fmt.Println(array6)
	fmt.Println(array7)
	fmt.Println()

	fmt.Printf("%-5T %d %v\n", array1, len(array1), array1)
	fmt.Printf("%-5T %d %v\n", array2, len(array2), array2)
	fmt.Printf("%-5T %d %v\n", array3, len(array3), array3)
	fmt.Printf("%-5T %d %v\n", array4, len(array4), array4)
	fmt.Printf("%-5T %d %v\n", array5, len(array5), array5)
	fmt.Printf("%-5T %d %v\n", array6, len(array6), array6)
	fmt.Printf("%-5T %d %v\n", array7, len(array7), array7)

	array8 := [5]int{1, 2, 3, 4, 5}
	array9 := [5]int{1, 2, 3, 4, 5}
	array10 := [...]string{"hello", "world"}

	fmt.Printf("%-5T %d %v\n", array8, len(array8), array8)
	fmt.Printf("%-5T %d %v\n", array9, len(array9), array9)
	fmt.Printf("%-5T %d %v\n", array10, len(array10), array10)
}
