package main

import "fmt"

type account struct {
	number   string
	balance  float64
	interest float64
}

func (a account) calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	// 구조체
	// 필드는 갖지만 메소드는 갖지 않는다
	// 객체지향 방식을 지원 -> 리시버를 통해 메소드랑 연결
	// 상속, 객체, 클래스 개념 없음
	// 구조체 -> 구조체 선언 -> 구조체 인스턴스(리시버)

	kim := account{number: "123-123", balance: 10000000, interest: 0.02}
	lee := account{number: "123-124", balance: 10000000}
	park := account{number: "123-125", interest: 0.012}
	cho := account{"123-126", 20000000, 0.012}

	fmt.Println("kim : ", kim)
	fmt.Println("lee : ", lee)
	fmt.Println("park : ", park)
	fmt.Println("cho : ", cho)

	fmt.Println("kim calculate : ", int(kim.calculate()))
}
