package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// 파일쓰기
	// CSV 파일 쓰기 예제
	// 패키지 저장소를 통해서 Excel 등 다양한 파일 형식 쓰기, 읽기 가능
	// bufio : 파일 용량이 클 경우 버퍼 사용 권장

	// 파일 생성
	file, error := os.Create("test_write.csv")
	errorCheck1(error)

	// 리소스 해제
	defer file.Close()

	// csv writer 생성
	//writer := csv.NewWriter(file)
	writer := csv.NewWriter(bufio.NewWriter(file))

	// csv 내용 쓰기
	writer.Write([]string{"kim", "4.8"})
	writer.Write([]string{"lee", "4.4"})
	writer.Write([]string{"park", "4.1"})
	writer.Write([]string{"hong", "3.2"})

	writer.Flush() // 버퍼 -> 파일로 쓰기

	fi, error := file.Stat()
	errorCheck2(error)

	fmt.Printf("csv 쓰기 파일 작업 후 파일 크기 (%d byte)\n", fi.Size())
	fmt.Println("csv 파일명 : ", fi.Name())
	fmt.Println("운영체제 파일 권한 : ", fi.Mode())
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
