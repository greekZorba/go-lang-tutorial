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

// 메서드 오버라이딩
func (e Executives) Calculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

func main() {
	// 구조체 임베디드 메서드 오버라이딩
	employee1 := Employee{"oh", 3000000, 2000000}
	employee2 := Employee{"mi", 4000000, 1000000}

	executive1 := Executives{
		Employee{"choi", 6500000, 3000000},
		2000000,
	}

	fmt.Println("employee1 : ", employee1)
	fmt.Println("employee2 : ", employee2)

	// Employee 구조체 통해서 메소드 호출
	fmt.Println("executive1 : ", int(executive1.Calculate()))
	fmt.Println("employee of executive : ", int(executive1.Employee.Calculate()))


}
