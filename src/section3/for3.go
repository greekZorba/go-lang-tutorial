package main

import "fmt"

func main() {

	// Loop1에 대한 정의까지 break 걸림
	// Loop와 관련 없는 코드가 들어가면 컴파일 에러 발생
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 1 {
				break Loop1
			}
			fmt.Println("example :", i, j)
		}
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("홀수 : ", i)
	}


}