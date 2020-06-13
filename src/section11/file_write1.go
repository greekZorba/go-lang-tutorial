package main

import (
	"fmt"
	"os"
)

func main() {
	// 파일 쓰기
	// Create : 새 파일 작성 및 파일 열기
	// Close : 리소스 닫기
	// Write, WriteString, WriteAt : 파일 쓰기
	// 각 운영체제 권한 주의(오류 메시지 확인)
	// 예외 처리 중요함

	file, error := os.Create("text_write.txt")

	errorCheck1(error)
	errorCheck2(error)

	// 리소스 해제
	defer file.Close()

	s1 := []byte{115, 112, 104, 101, 111}
	n1, error := file.Write(s1)
	errorCheck2(error)

	fmt.Printf("쓰기 작업 완료 1 (%d byte) \n", n1)

	file.Sync() // Write Commit(Stable)

	s2 := "\nHello Golang!\n File writes Test! - 1\n"
	n2, error := file.WriteString(s2)
	errorCheck2(error)

	fmt.Printf("쓰기 작업 완료 2 (%d byte) \n", n2)

	file.Sync() // Write Commit(Stable)

	s3 := "Test WriteAt - 2\n"
	n3, error := file.WriteAt([]byte(s3), 100) // Len(offset) 조절하면서 테스트
	errorCheck2(error)

	fmt.Printf("쓰기 작업 완료 3 (%d byte) \n", n3)

	file.Sync()

	n4, error := file.WriteString("Hello Golang! \n File writes test - 3")
	errorCheck1(error)

	fmt.Printf("쓰기 작업 완료 4 (%d byte) \n", n4)
}

func errorCheck1(e error) {
	if e != nil {
		panic(e)
	}

}

func errorCheck2(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}

}
