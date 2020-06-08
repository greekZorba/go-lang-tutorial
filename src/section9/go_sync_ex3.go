package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	/*
	고루틴 동기화 고급
	원자성 사용 -> 기능적으로 분할 불가능한 완전 보증된 일련의 조작, 모두 성공하거나 실패
	모든 조작이 완료될 때까지 다른 프로세스 개입 불가
	sync/atomic에서 원자적 연산자 제공
	https://golang.org/pkg/sync/atomic 에서 계열 확인 가능
	주로 공영변수에 관한 계산 사용
	*/

	//원자성 사용 안할 경우 예제
	runtime.GOMAXPROCS(runtime.NumCPU())

	var cnt int64 = 0
	waitGroup := new(sync.WaitGroup)

	for i := 1; i <= 5000; i++ {
		waitGroup.Add(1)
		go func(n int) {
			cnt++
			waitGroup.Done()
		}(i)
	}

	for i := 1; i <= 2000; i++ {
		waitGroup.Add(1)
		go func(n int) {
			cnt--
			waitGroup.Done()
		}(i)
	}

	// Add == Done 횟수 같아야함
	waitGroup.Wait()
	fmt.Println("WaitGroup end!! cnt : ", cnt)
}
