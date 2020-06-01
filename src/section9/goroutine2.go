package main

import (
	"fmt"
	"time"
)

func example(name string) {
	fmt.Println(name, " starts : ", time.Now())
	for i := 0; i < 1000; i++ {
		fmt.Println(name, ">>>>>>", i)
	}
	fmt.Println(name, " ends : ", time.Now())
}

func main() {
	example("t1")
	fmt.Println("Main Routine starts !! ", time.Now())
	go example("t2")
	go example("t3")
	time.Sleep(4 * time.Second)
	fmt.Println("Main Routine ends !! ", time.Now())
}
