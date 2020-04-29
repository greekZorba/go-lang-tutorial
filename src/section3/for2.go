package main

import "fmt"

func main() {

	sum := 0
	for i := 0; i < 100; i++ {
		sum += 1
	}

	fmt.Println("sum :", sum)

	sum2, i := 0, 0
	for i <= 100 {
		sum2 += i
		i++
		// j := i++ -> Go에서 후치연산은 반환값 X
	}

	fmt.Println("sum2 :", sum2)

	sum3, i := 0, 0
	for { // while 형태랑 비슷
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}

	fmt.Println("sum3 :", sum3)

	for i, j := 0, 0; i < 10; i, j = i+1, j+10 {
		fmt.Println("i :", i, "j :", j)
	}
}
