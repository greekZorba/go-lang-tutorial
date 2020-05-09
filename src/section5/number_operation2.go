package main

import "fmt"

func main() {

	var n1 = 100
	var n2 = 10

	fmt.Println(n1 + n2)
	fmt.Println(n1 * n2)
	fmt.Println(n1 - n2)
	fmt.Println(n1 / n2)
	fmt.Println(n1 % n2)
	fmt.Println(n1 << 2)
	fmt.Println(n1 >> 2)
	fmt.Println(^n1)

	var n3 int = 2
	var n4 float32 = 8.2

	fmt.Println(n3 + int(n4))
	fmt.Println(float32(n3) + n4)
}
