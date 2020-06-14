package main

import (
	"bufio"
	_ "bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// CSV 파일 읽기 예제

	// 파일 생성
	file, error := os.Open("sample.csv")
	errorCheck2(error)

	// 리소스 해제
	defer file.Close()

	// csv reader 생성
	reader := csv.NewReader(bufio.NewReader(file)) // 권장

	row, error := reader.Read() // 1개의 row 단위로 읽기
	errorCheck1(error)

	row2, error := reader.Read()
	fmt.Println("CSV Read example")
	fmt.Println(row)
	fmt.Println(row[0], row[1:5])
	fmt.Println(row2)
	fmt.Println("=========================")

	rows, error := reader.ReadAll() // 전체 row 읽기
	errorCheck2(error)
	fmt.Println("CSV ReadAll example")
	fmt.Println(rows[0])
	fmt.Println(rows[5][1])
	fmt.Println("=========================")

	fileInfo, error := file.Stat()
	errorCheck2(error)

	fmt.Println("파일 정보 출력 : ", fileInfo)
	fmt.Println("파일 이름 : ", fileInfo.Name())
	fmt.Println("파일 크기 : ", fileInfo.Size())
	fmt.Println("파일 수정 시간 : ", fileInfo.ModTime())

	// Row 단위로 CSV 파일 읽기
	for i, row := range rows {
		for j := range row {
			fmt.Printf("%s	", rows[i][j])
		}
	}
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
