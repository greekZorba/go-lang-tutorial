package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) calculate() float64 {
	return a.balance + (a.balance * a.interest)
}


func main() {
	// 다양한 선언 방법
	// &struct, struct : &struct 포인터를 받아오기 역참조를 또 하기 때문에 속도는 조금 느리다
	// 인터페이스 메서드를 선언만 해둔 후 -> 오버라이딩해서 메서드에 포인트 리시버를 사용할 경우 반드시 &struct
	//

	// 선언 방법1
	var kim *Account = new(Account)
	kim.number = "123-123"
	kim.balance = 10000000
	kim.interest = 0.12

	// 선언 방법2
	hong := &Account{number: "123-124", balance: 10000000, interest: 4.2}

	// 선언 방법3
	lee := new(Account)
	lee.number = "123-125"
	lee.balance = 30000000
	lee.interest = 0.15

	fmt.Println("kim : ", kim)
	fmt.Println("hong : ", hong)
	fmt.Println("lee : ", lee)

	fmt.Printf("kim : %#v \n", kim)
	fmt.Printf("hong : %#v \n", hong)
	fmt.Printf("lee : %#v \n", lee)

	fmt.Println("kim : ", int(kim.calculate()))
	fmt.Println("hong : ", int(hong.calculate()))
	fmt.Println("lee : ", int(lee.calculate()))
}
