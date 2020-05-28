package main

import "fmt"

type Employee struct {
	name string
	salary float64
	bonus float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

type Executives struct {
	Employee // is a 관계
	specialBonus float64
}

func main() {
	// 구조체 임베디드 패턴
	// 다른 관점으로 메소드를 재사용하는 장점 제공
	// 상속을 허용하지 않는 Go 언어에서 메소드 상속 활용을 위한 패턴

	employee1 := Employee{"kim", 3000000, 2000000}
	employee2 := Employee{"park", 4000000, 1000000}

	executive1 := Executives{
		Employee{"choi", 6500000, 3000000},
		2000000,
	}

	fmt.Println("employee1 : ", employee1)
	fmt.Println("employee2 : ", employee2)

	// Employee 구조체 통해서 메소드 호출
	fmt.Println("executive1 : ", int(executive1.Calculate() + executive1.specialBonus))


}
