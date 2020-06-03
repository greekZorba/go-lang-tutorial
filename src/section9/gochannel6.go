package main

import "fmt"

func main() {
	// Close : 채널 닫기

	ch := make(chan string)

	go func() {
		for i := 0; i < 3 ; i++ {
			ch <- "study Go!"
		}
	}()

	val1, ok1 := <- ch
	val2, ok2 := <- ch
	val3, ok3 := <- ch

	fmt.Println("val1 : ", val1, "ok1 : ", ok1)
	fmt.Println("val2 : ", val2, "ok2 : ", ok2)
	fmt.Println("val3 : ", val3, "ok3 : ", ok3)

	close(ch)
	val4, ok4 := <- ch
	fmt.Println("val4 : ", val4, "ok4 : ", ok4)
}
