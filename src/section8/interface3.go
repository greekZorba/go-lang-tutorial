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

func (d Dog) bite() {
	fmt.Println(d.name, " Dog bites")
}

func (d Dog) sounds() {
	fmt.Println(d.name, " Dog barks")
}

func (d Dog) run() {
	fmt.Println(d.name, " Dog is running")
}

func (c Cat) bite() {
	fmt.Println(c.name, " Cat bites")
}

func (c Cat) sounds() {
	fmt.Println(c.name, " Cat cries")
}

func (c Cat) run() {
	fmt.Println(c.name, " Cat is running")
}

// 동물의 행동 인터페이스 선언
type Behavior interface {
	bite()
	sounds()
	run()
}

// 인터페이스를 파라미터로 받는다
func act(animal Behavior) {
	animal.bite()
	animal.sounds()
	animal.run()
}

func main() {
	// 인터페이스 규격화 역할 이해
	// 인터페이스에 정의된 메서드 사용 유도
	// 코드의 가독성 및 유지보수 증가

	// 덕타이핑 : 구조체 및 변수의 값이나 타입은 상관하지 않고 오로지 구현된 메서드로만 판단하는 방식
	// Go 덕타이핑의 중요한 특징 : 오리처럼 걷고, 소리 내고, 헤엄 등 행동이 같으면 오리라고 볼 수 있다 -> 의미
	dog1 := Dog{"songYee", 10}
	cat1 := Cat{"myo", 6}

	// 개 행동 실행
	act(dog1)

	// 고양이 행동 실행
	act(cat1)
}
