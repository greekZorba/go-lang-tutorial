package main

import (
	"fmt"
	"math"
)

func Power(f float64, i float64) (float64, error) {
	if f == 0 {
		return 0, fmt.Errorf("%g은 사용불가능 합니다", f)
	}

	return math.Pow(f, i), nil
}

func main() {

	if zeroInput, error := Power(0, 4); error != nil {
		fmt.Println("error : ", error.Error())
	} else {
		fmt.Println("zeroInput : ", zeroInput)
	}

	if someInput, error := Power(4, 3); error != nil {
		fmt.Println("error : ", error.Error())
	} else {
		fmt.Println("someInput : ", someInput)
	}
}
