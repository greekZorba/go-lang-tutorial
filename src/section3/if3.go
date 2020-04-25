package main

import "fmt"

func main() {
	i := 100

	if i > 120 {
		fmt.Println("i는 120보다 크다")
	} else if i <= 120 && i > 100 {
		fmt.Println("i는 120보다 작거나 같고 100보다 크다")
	} else {
		fmt.Println("i는 100보다 작거나 같다")
	}
}
