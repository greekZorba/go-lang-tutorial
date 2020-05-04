package main

import (
	"fmt"
	"os"
)
func main() {
	// 패키지 : 코드 구조화(모듈) 및 재사용
	// 응집도는 높히고, 결합도는 낮춘다
	// Go는 패키지 단위의 독립이고 작은 단위로 개발 -> 작은 패키지를 결합해서 프로그램을 작할 것을 권고
	// 패키지 이름 = 디렉토리 이름
	// 같은 패키지 내 -> 소스 파일들은 디렉토리명을 패키지명으로 사용한다
	// 네이밍 규칙 : 소문자 private, 대문자 public
	// 메인 패키지는 특별하게 인식함 -> 컴파일러 공유 라이브러리 x, 프로그램의 start point

	var name string

	fmt.Println("이름은? ")

	fmt.Scanf("%s", &name)

	fmt.Fprintf(os.Stdout, "Hi ! %s \n", name)
}
