package main

import "fmt"

func sayHello(message string) {
	defer func() {
		fmt.Println(message)
	}()

	func() {
		fmt.Println("펭하!!")
	}()
}

func main() {

	sayHello("defer function start!!")
}
