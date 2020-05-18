package main

import "fmt"

func main() {
	// 포인터
	// Go : 포인터 지원
	// 변수의 지역성, 연속된 메모리 참조, 힙, 스택
	// 파이선, 자바(JRE) -> 컴파일러, 인터프리터
	// 포인터 지원x(파이선, C#, Java 등)
	// 주소의 값은 직접 변경 불가능(잘못된 코딩으로 인한 버그 방지)
	// *(애스터리스크)로 사용
	// nil로 초기화(nil == 0)

	var a *int
	var b *int = new(int)

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)

	i := 7

	fmt.Println("i : ", i, " address : ", &i)

	a = &i // 주소값 전달
	b = &i

	fmt.Println("a : ", a)
	fmt.Println("a : ", &a)
	fmt.Println("a : ", *a) // 역참조

	fmt.Println("b : ", b)
	fmt.Println("b : ", &b)
	fmt.Println("b : ", *b) // 역참조

	c := &i
	d := &i

	*d = 10

	fmt.Println("c : ", c)
	fmt.Println("c : ", &c)
	fmt.Println("c : ", *c)

	fmt.Println("d : ", d)
	fmt.Println("d : ", &d)
	fmt.Println("d : ", *d)

}
