package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
func main() {
	// 파일 읽기, 버퍼 사용 쓰기 -> bufio 패키지 활용
	// ioutil, bufio 등은 io.Reader, io.Writer 인터페이스를 구현함 -> 즉, Writer, Read 메서드를 사용 가능
	/*
		type Reader interface {
			Read(p []byte) (n int, err error)
		}

		type Writer interface {
			Write(p []byte) (n int, err error)
		}
	*/
	// 즉, bufio의 NewReader, NewWriter를 통해서 객체 반환 후 메서드 사용

	// bufio(buffered io) 패키지
	// https://golang.org/pkg/bufio

	// 파일 열기
	// 두번째 매개변수 확인
	// https://golang.org/pkg/os/#pkg-constants

	// bufio 파일 쓰기 예제
	file, error := os.OpenFile("test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	errorCheck(error)

	writer := bufio.NewWriter(file) // Writer 반환(버퍼 사용)
	writer.WriteString("Hello Golang\n File Writes Test1\n")
	writer.Write([]byte("Hello Golang\n File Writes Test2\n"))

	// 버퍼 정보 출력
	fmt.Printf("사용한 Buffer Size (%d byte)\n", writer.Buffered())
	fmt.Printf("남은 Buffer Size (%d byte)\n", writer.Available())
	fmt.Printf("전체 Buffer Size (%d byte)\n", writer.Size())

	writer.Flush() // 버퍼를 비우고 반영(버퍼에 내용을 디스크에 저장)

	// Reader 반환
	reader := bufio.NewReader(file)
	fi, error := file.Stat()
	errorCheck(error)

	b := make([]byte, fi.Size())

	fmt.Println("파일 정보 출력 : ", fi)
	fmt.Println("파일 이름 : ", fi.Name())
	fmt.Println("파일 크기 : ", fi.Size())
	fmt.Println("파일 수정 시간 : ", fi.ModTime())

	fmt.Println("========================")

	file.Seek(0, io.SeekStart)
	data, _ := reader.Read(b) // 읽기(ReadLine, ReadByte, ReadBytes 등)

	fmt.Printf("전체 Buffer Size (%d byte)\n", reader.Size())
	fmt.Printf("읽기 작업 완료 (%d byte)\n", data)
	fmt.Println(string(b))

	defer file.Close()
}

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}
