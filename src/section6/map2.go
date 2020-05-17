package main

import "fmt"

func main() {
	// 맵 조회 및 순회

	map1 := map[string]string{
		"daum" : "daum.com",
		"naver" : "naver.com",
		"google" : "google.com",
	}

	fmt.Println("map1 : ", map1)
	fmt.Println("daum url : ", map1["daum"])

	// 순서 없으므로 랜덤
	for index, value := range map1 {
		fmt.Println("index : ", index, "value : ", value)
	}

	for _, value := range map1 {
		fmt.Println("value : ", value)
	}

}
