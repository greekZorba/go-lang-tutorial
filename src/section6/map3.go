package main

import "fmt"

func main() {
	// 맵 변경 및 삭제
	map1 := map[string]string{
		"daum" : "daum.com",
		"naver" : "naver.com",
		"google" : "google.com",
	}

	// 추가
	map1["wechat"] = "wechat.com"

	// 수정
	map1["daum"] = "modifiedDaum.com"

	// 삭제
	delete(map1, "google")

	fmt.Println("map1 : ", map1)
}
