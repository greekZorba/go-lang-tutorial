package main

import "fmt"

func multiply(x, y int) (r1 int, r2 int) {
	r1 = x * 20
	r2 = y * 12

	return r1, r2
}

func main() {
	a, b := multiply(15, 32)
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
}
