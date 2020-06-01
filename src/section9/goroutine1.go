package main

import (
	"fmt"
	"time"
)

func example1() {
	fmt.Println("example1 starts!! ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("example1 ends!! ", time.Now())
}

func example2() {
	fmt.Println("example2 starts!! ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("example2 ends!! ", time.Now())
}

func example3() {
	fmt.Println("example3 starts!! ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("example3 ends!! ", time.Now())
}

func main() {
	// 고루틴(Goroutine)
	// 타 언어의 쓰레드(Thread)와 비슷한 기능
	// 생성 방법 매우 간단, 리소스 매우 적게 사용
	// 즉, 수많은 고루틴 동시 생성 실행 가능
	// 비동기적 함수 루틴 실행(적은 용량 차지) -> 채널을 통한 통신 가능
	// 공유 메모리 사용 시에 정확한 동기화 코딩 필요
	// 싱글 루틴에 비해 항상 처리 결과가 빠른 건 아니다

	/*
		멀티 쓰레드 장점 및 단점
		장점 : 응답성 향상, 자원 공유를 효율적으로 활용 사용, 작업 코드가 분리되어 코드 간결
		단점 : 구현하기 어려움, 테스트 및 디버깅 어려움, 전체 프로세스의 사이드 이펙트, 성능 저하, 동기화 코딩 반드시 숙지, 데드락..
	*/

	example1()
	fmt.Println("Main Routine starts ! ")
	go example2()
	go example3()
	time.Sleep(4 * time.Second)
	fmt.Println("Main Routine ends ! ")
}
