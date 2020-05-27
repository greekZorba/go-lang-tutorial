package main

import "fmt"

type car struct {
	color string
	name string
}

func main() {
	car1 := car{"red", "sonata"}
	car2 := new(car)
	car2.color, car2.name = "white", "avante"
	car3 := &car{}
	car4 := &car{"black", "genesis"}

	fmt.Println(car1)
	fmt.Println(car2)
	fmt.Println(car3)
	fmt.Println(car4)
}
