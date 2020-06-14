package main

import (
	_ "bufio"
	_ "encoding/csv"
	"fmt"
	"os"
)

func main() {
	// 파일 읽기
	// Open: 기존 파일 열기
	// Close: 리소스 닫기
	// Read, ReadAt: 파일 읽기
	// 각 운영체제 권한 주의(오류 메시지 확인)
	// 예외 처리 중요!
	// 탐색 Seek 중요

	file, error := os.Open("sample.txt")
	errorCheck1(error)

	fileInfo, error := file.Stat() // 파일 사이즈 확인 위해 정보 획득
	errorCheck2(error)

	fd1 := make([]byte, fileInfo.Size()) // 슬라이스에 읽는 내용을 담는다
	ct1, error := file.Read(fd1)

	fmt.Println("파일 정보 출력 : ", fileInfo)
	fmt.Println("파일 이름 : ", fileInfo.Name())
	fmt.Println("파일 크기 : ", fileInfo.Size())
	fmt.Println("파일 수정 시간 : ", fileInfo.ModTime())
	fmt.Printf("읽기 작업 완료 : (%d bytes) \n", ct1)
	//fmt.Println(fd1) // 타입 변환 없을경우
	fmt.Println(string(fd1)) // 타입 변환 한 경우

	fmt.Println("===============================")

	// 탐색 : Seek(Offset)
	// whence - 0: 처음 위치, 1: 현재 위치, 2: 마지막 위치
	o1, error := file.Seek(20, 0)
	errorCheck1(error)

	fd2 := make([]byte, 20)
	ct2, error := file.Read(fd2)
	errorCheck1(error)

	fmt.Printf("읽기 작업 완료 : (%d bytes) (%d ret)\n", ct2, o1)
	fmt.Println(string(fd2))
	fmt.Println("===============================")

	o2, error := file.Seek(0,0)
	errorCheck2(error)

	fd3 := make([]byte, 50)
	ct3, error := file.ReadAt(fd3, 8) // offset 위치부터 읽어온다.

	fmt.Printf("읽기 작업 완료 : (%d bytes) (%d ret)\n", ct3, o2)
	fmt.Println(string(fd3))
	fmt.Println("===============================")

	file.Close()
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
