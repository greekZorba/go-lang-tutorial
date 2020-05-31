package main

import "fmt"

func main() {
	// 실제 타입 검사 switch 사용
	// 빈 인터페이스는 어떠한 자료형도 전달 받을 수 있으므로, 타입 체크를 통해 형 변환 후 사용가능

	hello := "hello"
	var emptyInterface interface{}
	count := 9
	flag := true

	checkType(hello)
	checkType(emptyInterface)
	checkType(count)
	checkType(flag)
}

func checkType(arg interface{}) {
	// org.(type)을 통해서 현재 데이터형 반환
	switch arg.(type) {
	case bool:
		fmt.Println("This is bool : ", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("This is int : ", arg)
	case float64:
		fmt.Println("This is float : ", arg)
	case string:
		fmt.Println("This is string : ", arg)
	case nil:
		fmt.Println("This is nil : ", arg)
	default:
		fmt.Println("what is this type? : ", arg)
	}
}
