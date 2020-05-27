package main

import "fmt"

func main() {
	// 구조체 익명 선언 및 활용
	// type 구조체명 타입
	car1 := struct {
		color string
		name  string
	}{"white", "bmw"}

	fmt.Printf("car1 : %#v \n", car1)

	cars := []struct {
		color string
		name  string
	}{{"black", "morning"}, {"blue", "benz"}, {"yellow", "matiz"}}

	fmt.Printf("cars : %#v \n", cars)

	for index, value := range cars {
		fmt.Println("index : ", index)
		fmt.Println("value : ", value)
	}
}
