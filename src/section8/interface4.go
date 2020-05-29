package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

func (d Dog) run() {
	fmt.Println(d.name, " Dog is running")
}

func (c Cat) run() {
	fmt.Println(c.name, " Cat is running")
}

func act(animal interface{ run() }) {
	animal.run()
}

func main() {
	// 익명 인터페이스 사용 예제 (즉시 선언 후 바로 사용)
	dog1 := Dog{"songYee", 10}
	cat1 := Cat{"myo", 6}

	// 개 행동 실행
	act(dog1)

	// 고양이 행동 실행
	act(cat1)
}
