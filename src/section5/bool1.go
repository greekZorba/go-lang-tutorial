package main

import "fmt"

func main() {
	// Bool(Boolean) : 참, 거짓
	// 조건부 논리 연산자랑 주로 사용 : !, ||(or), &&(and)
	// 암묵적 형변환x : 0, Nil -> False 변환 없음

	var b1 bool = true
	var b2 bool = false
	b3 := true

	fmt.Println("b1 :", b1)
	fmt.Println("b2 :", b2)
	fmt.Println("b3 :", b3)

}
