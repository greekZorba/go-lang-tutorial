package main

import (
	"fmt"
	"log"
)

func notZero(n int) (string, error) { // 메서드 리턴 값 error 타입 중요!
	if n != 0 {
		s := fmt.Sprint("this is not zero -> n : ", n)
		return s, nil
	}

	return "this is zero", fmt.Errorf("%d을 입력했습니다. 에러발생!", n)
}

func main() {
	// Errorf를 이용한 에러 처리
	a, err := notZero(2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("a : ", a)

	b, err := notZero(0)

	if err != nil {
		log.Fatal(err)
	}

	// Fatal 이후 프로그램 종료 후 실행 중지
	fmt.Println("b : ", b)
	fmt.Println("End Error handling!")
}
