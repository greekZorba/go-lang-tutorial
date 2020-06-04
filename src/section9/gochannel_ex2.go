package main

import "fmt"

func sum(cnt int) <-chan int {
	c := make(chan int)

	var sum int
	go func() {
		for i := 1; i <= cnt; i++ {
			sum += i
		}
		c <- sum
	}()

	return c
}

func main() {
	// 채널
	// 채널 또한 함수의 반환값으로 사용가능

	c := sum(100)

	fmt.Println("c :", <-c)
}
