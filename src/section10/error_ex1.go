package main

import (
	"errors"
	"fmt"
	"math"
)

func Power(f float64, i float64) (float64, error) {
	if f == 0 {
		return 0, errors.New("0은 사용불가능 합니다")
	}

	return math.Pow(f, i), nil
}

func main() {

	if zeroInput, error := Power(0, 4); error != nil {
		fmt.Println("error : ", error)
	} else {
		fmt.Println("zeroInput : ", zeroInput)
	}

	if someInput, error := Power(4, 3); error != nil {
		fmt.Println("error : ", error)
	} else {
		fmt.Println("someInput : ", someInput)
	}
}
