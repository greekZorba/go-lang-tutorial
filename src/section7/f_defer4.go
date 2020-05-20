package main

import "fmt"

func start(t string) string {
	fmt.Println("start : ", t)
	return t
}

func end(t string) string {
	fmt.Println("end : ", t)
	return t
}

func a() {
	defer end(start("b")) // argument 로 들어간 함수는 defer 영향을 받지 않고 바로 실행된다
	fmt.Println("in a")
}

func main() {
	a()
}
