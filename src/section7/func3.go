package main

import "fmt"

func multiply(x, y int) (int, int) {
	return x * 10, y * 10
}

// 다중값 반환
func array_multiply(a, b, c, d, e int) (int, int, int, int, int) {
	return a * 10, b * 10, c * 10, d * 10, e * 10
}

func main() {
	a, b := multiply(10, 18)
	c, _ := multiply(1, 5)
	_, d := multiply(2, 9)

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	fmt.Println("d : ", d)

	e, f, g, h, i := array_multiply(10, 5, 212, 424, 12)

	fmt.Println("e : ", e)
	fmt.Println("f : ", f)
	fmt.Println("g : ", g)
	fmt.Println("h : ", h)
	fmt.Println("i : ", i)

}
