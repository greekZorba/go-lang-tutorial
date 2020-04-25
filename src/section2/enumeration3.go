package main

import "fmt"

func main() {
	// "_"는 iota를 스킵처리 함
	const (
		_ = iota
		A
		_
		C
	)

	const (
		_ = iota + 0.75 * 2
		DEFAULT
		SILVER
		_
		PLATINUM
	)

	fmt.Println(A, C)
	fmt.Println(DEFAULT, SILVER, PLATINUM)

}
