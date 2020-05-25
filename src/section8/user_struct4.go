package main

import "fmt"

type shoppingBasket struct {
	cnt, price int
}

func (b shoppingBasket) purchase() int {
	return b.cnt * b.price
}

// 원본 수정됨(참조 전달형식)
func (b *shoppingBasket) rePurchase(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

// 원본 수정 안됨(값 전달형식)
func (b shoppingBasket) rePurchase2(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func main() {
	// 리시버(구조체 메서드) 전달(값, 참조) 형식
	// 함수는 기본적으로 값 호출 -> 변수의 값이 복사 후 내부 전달(원본 수정X) -> 맵, 슬라이스 등은 참조 전달
	// 리시버(구조체)도 마찬가지로 포인터를 활용해서 메소드 내에서 원본 수정 가능
	basket := shoppingBasket{3, 5000}
	fmt.Println("basket totalPrice : ", basket.purchase())

	// 참조 전달 (원본값 수정)
	basket.rePurchase(7, 5000)
	fmt.Println("basket rePurchase : ", basket.purchase())

	// 값 전달 (원본값 수정X)
	basket.rePurchase2(20, 10000)
	fmt.Println("basket rePurchase2 : ", basket.purchase())

}
