package main

import (
	"fmt"
	"runtime"
)

type count struct {
	num int
}

func (c *count) increment() {
	c.num++
}

func (c *count) result() int {
	return c.num
}

func main() {
	/*
		고루틴 동기화 예제
		실행흐름 제어 및 변수 동기화 가능
		공유 데이터 보호가 중요
	*/

	// 동기화 사용하지 않는 경우 예제
	// 시스템 전체 CPU 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count{num: 20}
	done := make(chan bool)

	for i := 1; i < 10000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() // CPU 양보
		}()
	}

	for i := 1; i < 10000; i++ {
		go func() {
			<-done
		}()
	}

	fmt.Println("result : ", c.result())
}
