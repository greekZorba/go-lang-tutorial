package main

import "fmt"

func main() {
	// 맵
	// 해쉬테이블, 딕셔너리, key-value로 저장
	// 레퍼런스 타입(참조값 전달)
	// 비교 연산자 사용 불가(참조 타입이므로)
	// 특징 : 참조 타입(key) 사용 불가능, 값(value)으로 모든 타입 사용 가능
	// make 함수 및 축약(리터럴)로 초기화 가능
	// 순서 없음

	var map1 map[string]int = make(map[string]int) // 정석
	var map2 = make(map[string]int)
	map3 := make(map[string]int)

	fmt.Println("map1 : ", map1)
	fmt.Println("map2 : ", map2)
	fmt.Println("map3 : ", map3)

	map4 := map[string]int{} // json 형태
	map4["apple"] = 25
	map4["banana"] = 44
	map4["berry"] = 11

	fmt.Println("map4 : ", map4)

	map5 := map[string]int{
		"apple" : 12,
		"banana" : 30,
		"berry" : 14,
	}

	fmt.Println("map5 : ", map5)

	map6 := make(map[string]int, 10)
	map6["apple"] = 12
	map6["banana"] = 122
	map6["berry"] = 121

	fmt.Println("map6 : ", map6)
}
