package main

import (
	"errors"
	"fmt"
	"log"
)

func notZero(n int) (string, error) { // 메서드 리턴 값 error 타입 중요!
	if n != 0 {
		s := fmt.Sprint("this is not zero -> n : ", n)
		return s, nil
	}

	return "this is zero", errors.New("0을 입력했습니다. 에러발생!") // 리턴값으로 errors.New 사용
}

func main() {
	// errors.New를 이용한 에러 처리
	a, err := notZero(2)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("a : ", a)

	b, err := notZero(0)

	if err != nil {
		log.Fatal(err.Error())
	}

	// Fatal 이후 프로그램 종료 후 실행 중지
	fmt.Println("b : ", b)
	fmt.Println("End Error handling!")
}
