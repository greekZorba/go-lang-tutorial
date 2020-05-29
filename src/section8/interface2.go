package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

func (d Dog) bite() {
	fmt.Println(d.name, "bites")
}

// 동물의 행동 인터페이스 선언
type Behavior interface {
	bite()
}

func main() {
	dog1 := Dog{"songYee", 10}

	var interface1 Behavior
	interface1 = dog1
	interface1.bite()

	dog2 := Dog{"arong", 5}
	interface2 := Behavior(dog2)
	interface2.bite()

	interfaces := []Behavior{dog1, dog2}
	for _, value := range interfaces {
		value.bite()
	}

	for index, _ := range interfaces {
		interfaces[index].bite()
	}
}
