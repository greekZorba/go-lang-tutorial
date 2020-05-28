package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func CalculateD(a Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CalculateP(a *Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func main()  {

	kim := Account{number: "123-412", balance: 10000, interest: 0.12}
	lee := Account{number: "533-412", balance: 40000, interest: 0.72}

	CalculateD(kim)
	CalculateP(&lee)

	fmt.Println("kim: ", kim)
	fmt.Println("lee: ", lee)
}