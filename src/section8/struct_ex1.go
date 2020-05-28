package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func NewAccount(number string, balance float64, interest float64) *Account { // 포인터 반환 아닌 경우는 값 복사
	return &Account{number, balance, interest} // 구조체 인스턴스를 생성한 뒤 포인터를 리턴
}

func main() {
	// 구조체 생성자 패턴

	kim := Account{number: "123-412", balance: 10000, interest: 0.12}

	var park *Account = new(Account)
	park.number = "312-124"
	park.balance = 20000
	park.interest = 1.24

	choi := NewAccount("123-521", 20000, 8.12)

	fmt.Println("kim : ", kim)
	fmt.Println("park : ", park)
	fmt.Println("choi : ", choi)
}
