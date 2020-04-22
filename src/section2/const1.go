package main

import "fmt"

func main() {
	// 상수
	// const 사용 초기화, 한 번 선언 후 변경금지, 고정된 값 관리용
	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10

	// 함수로 값을 가져오는건 컴파일 에러남. 값이 매번 바뀔 수 있어서
	// const d = getHeight()
	const e = 12.6
	const f = true
	 // 에러 발생
	 /*
	 	const g string
	 	g = "Test3"
	 */
	 fmt.Println("a : ", a, "b : ", b, "c : ", c, "e : ", e, "f : ", f)
}
