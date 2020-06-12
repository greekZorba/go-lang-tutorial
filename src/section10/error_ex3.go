package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

// 예외(에러) 처리 구조체
type PowError struct {
	time    time.Time // 에러 발생 시간
	value   float64   // 파라미터
	message string    // 에러 메시지
}

func (e PowError) Error() string {
	return fmt.Sprintf("[%v] Error - Input Value(value: %g) - %s", e.time, e.value, e.message)
}

func Power(f float64, i float64) (float64, error) {
	if f == 0 {
		return 0, PowError{time: time.Now(), value: f, message: "0은 사용할 수 없습니다"}
	}

	if math.IsNaN(f) {
		return 0, PowError{time: time.Now(), value: f, message: "f가 숫자가 아닙니다"}
	}

	if math.IsNaN(i) {
		return 0, PowError{time: time.Now(), value: i, message: "i가 숫자가 아닙니다"}
	}

	return math.Pow(f, i), nil
}

func main() {
	// error 타입이 아닌 경우 에러 처리 방법
	// Error 메서드를 구현해서 사용자 정의 에러 처리 예제 심화
	// 구조체를 사용해서 세부적인 정보 출력

	value1, error1 := Power(2, 4)
	if error1 != nil {
		log.Fatal(error1)
	}

	fmt.Println("value1 : ", value1)

	value2, error2 := Power(0, 4)
	if error2 != nil {
		//log.Fatal(error2)
		fmt.Println("time : ", error2.(PowError).time)
		fmt.Println("value : ", error2.(PowError).value)
		fmt.Println("message : ", error2.(PowError).message)
	}

	fmt.Println("value2 : ", value2)
}
