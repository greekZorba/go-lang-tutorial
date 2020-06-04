package main

import "fmt"

func receiveOnly(c int) <-chan int {
	tot := make(chan int)

	var sum int
	go func() {
		for i := 1; i <= c; i++ {
			sum += i
		}
		tot <- sum
		tot <- 2500
		tot <- 1000
		close(tot)
	}()

	return tot
}

func total(c <-chan int) <-chan int {
	tot := make(chan int)

	go func() {
		a := 0
		for i := range c {
			a += i
		}
		tot <- a
	}()

	return tot
}

func main() {

	c := receiveOnly(100) // 채널 반환
	fmt.Println("c : ", <-c)
	output := total(c) // 채널 전달 후 반환

	fmt.Println("output : ", <-output)
}
