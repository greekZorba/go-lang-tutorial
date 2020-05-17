package main

import "fmt"

func main() {
	// 맵 조회할 경우 주의할 점

	map1 := map[string]string{
		"daum":   "daum.com",
		"naver":  "naver.com",
		"google": "google.com",
	}

	// 두 번째 변수를 선언해서 키 존재 유무를 확인할 수 있음
	daumUrl, isExist := map1["daum"]
	facebookUrl := map1["facebook"]
	instagramUrl, ok := map1["instagram"]

	fmt.Println("daumUrl : ", daumUrl)
	fmt.Println("isExistDaumUrl : ", isExist)
	fmt.Println("facebookUrl : ", facebookUrl)
	fmt.Println("instagramUrl : ", instagramUrl)
	fmt.Println("isExistInstagramUrl : ", ok)

	if value, isExist := map1["naver"]; isExist {
		fmt.Println("value : ", value)
	} else {
		fmt.Println("value is not exist")
	}

	if _, isExist := map1["instagram"]; !isExist {
		fmt.Println("instagram is not exist")
	}
}
