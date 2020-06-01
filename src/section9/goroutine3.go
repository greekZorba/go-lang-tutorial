package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func example(name int) {
	r := rand.Intn(100)

	fmt.Println(name, " starts ! ", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, " >>>>> ", r, i )
	}
	fmt.Println(name, " ends ! ", time.Now())
}

func main() {
	// 멀티코어(다중 CPU) 최대한 활용

	runtime.GOMAXPROCS(runtime.NumCPU()) // 현 시스템의 CPU 코어 개수 반환 후 설정
	fmt.Println("current my cpu : ", runtime.GOMAXPROCS(0)) // 설정값 출력

	fmt.Println("Main routine starts : ", time.Now())

	for i := 0; i < 100; i++ {
		go example(i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Main routine ends : ", time.Now())

}
