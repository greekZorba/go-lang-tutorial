package main

import "fmt"

func main() {
	 const a, b int = 1, 2
	 const c, d, e = true, 0.12, "hello!"
	 const (
	 	x, y int16 = 50, 90
	 	i, k = "DATA", 771
	 )

	 fmt.Println(a, b, c, d, e, x, y, i, k)
}
