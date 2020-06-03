package main

import (
	"fmt"
	"time"
)

func main() {
	// 동기 : 버퍼 미사용
	ch := make(chan bool)
	cnt := 6

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("func : ", i)
			time.Sleep(3 * time.Second)
		}
	}()

	for i := 0; i < cnt; i++ {
		<- ch
		fmt.Println("Main : ", i)
	}
}
