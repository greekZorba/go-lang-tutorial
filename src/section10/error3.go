package main

import (
	"errors"
	"fmt"
)

func main() {
	// errors 패키지의 New 메서드를 활용한 에러 생성
	// Errorf보다 더 많이 사용

	var err1 error = errors.New("Error occured - 1")
	err2 := errors.New("Error occured - 2")

	fmt.Println("error1 : ", err1)
	fmt.Println("error1.Error() : ", err1.Error())

	fmt.Println("error2 : ", err2)
	fmt.Println("error2.Error() : ", err2.Error())
}
