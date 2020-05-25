package main

import "fmt"

type totalCost func(int, int) int

func describe(cnt int, price int, fn totalCost) {
	fmt.Printf("cnt : %d, price : %d, totalCost : %d", cnt, price, fn(cnt, price))
}

func main() {
	// 함수 사용자 정의 타입
	var orderPrice totalCost
	orderPrice = func(cnt int, price int) int{
		return (cnt * price) + 10000
	}

	fmt.Println("orderPrice : ", orderPrice(10, 200))

	describe(200, 300, orderPrice)
}
