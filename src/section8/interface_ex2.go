package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 타입 변환
	// 실행(런타임) 시에는 인터페이스에 할당한 변수는 실제 타입으로 변환 후 사용해야 하는 경우
	// 인터페이스.(타입) -> 형변환
	// interfaceVal.(type)

	var a interface{} = 15
	b := a
	c := a.(int)
	//d := a.(float64)

	fmt.Println("a : ", a)
	fmt.Println("a : ", reflect.TypeOf(a))
	fmt.Println("b : ", b)
	fmt.Println("b : ", reflect.TypeOf(b))
	fmt.Println("c : ", c)
	fmt.Println("c : ", reflect.TypeOf(c))
	//fmt.Println("d : ", d)

	// 저장된 실제 타입 검사
	if v, ok := a.(int); ok {
		fmt.Println("v : ", v, "ok : ", ok)
	}
}
