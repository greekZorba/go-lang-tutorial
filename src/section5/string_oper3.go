package main

import (
	"fmt"
	"strings"
)

func main() {
	// 결합 : 일반 연산
	var str1 = "The Go programming language is an open source project to make programmers more productive." +
		"The Go programming language is an open source project to make programmers more productive." +
		"The Go programming language is an open source project to make programmers more productive." +
		"The Go programming language is an open source project to make programmers more productive." +
		"The Go programming language is an open source project to make programmers more productive."

	var str2 = "Go is expressive, concise, clean, and efficient. " +
		"Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, " +
		"while its novel type system enables flexible and modular program construction. " +
		"Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. " +
		"It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language."

	fmt.Println(str1 + str2)

	// 결합 : join
	var strSet []string // slice 선언
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)

	fmt.Println(strings.Join(strSet, "-----------"))
}
