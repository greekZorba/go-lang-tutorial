package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	/*
	고루틴 동기화 객체
	동기화 상태(조건) 메서드 사용
	Wait, notify, notifyAll: 기타 언어
	Wait, Signal, Broadcast
	 */

	// 시스템 전체 CPU 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	mutex := new(sync.Mutex)
	condition := sync.NewCond(mutex)

	c := make(chan int, 5) // 비동기 버퍼 채널

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("Goroutine Waiting : ", n)
			condition.Wait() // 고루틴 대기(멈춤)
			fmt.Println("Waiting end : ", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println("c : ", <-c)
	}

	// 한개씩 깨움(모든 고루틴 생성 후)
	for i := 0; i < 5; i++ {
		mutex.Lock()
		fmt.Println("Wake goroutine(Signal) : ", i)
		condition.Signal()
		mutex.Unlock()
	}

	//한번에 다 깨움
	//mutex.Lock()
	//fmt.Println("Wake Goroutine(Broadcast)")
	//condition.Broadcast()
	//mutex.Unlock()

	time.Sleep(2 * time.Second)
}
