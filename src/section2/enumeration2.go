package main

import "fmt"

func main() {
	const (
		A = iota + 1
		B
		C
	)

	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
		May
		Jun
	)

	fmt.Println(A, B, C)
	fmt.Println(Jan, Feb, Mar, Apr, May, Jun)
}
