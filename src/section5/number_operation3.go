package main

import (
	"fmt"
	"math"
)

func main() {
	// overFlow
	var n1 uint8 = math.MaxUint8 + 1
	fmt.Println(n1 )
}
