package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) CalculateD(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a *Account) CalculateP(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func main()  {
	// 구조체 인스턴스 값 변경시 -> 포인터 전달, 보통의 경우 -> 값 전달
	kim := Account{number: "123-412", balance: 10000, interest: 0.12}
	lee := Account{number: "533-412", balance: 40000, interest: 0.72}

	kim.CalculateD(4000000)
	lee.CalculateP(5000000)

	fmt.Println("kim: ", int(kim.balance))
	fmt.Println("lee: ", int(lee.balance))
}