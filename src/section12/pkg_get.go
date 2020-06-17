// 사용자 패키지 설치 및 활용 예제
package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	// 외부 저장소 패키지 설치
	// 2가지 방법
	// 1. import 선언 후 폴더 이동 후 go get 설치 -> 사용
	// 2. go get 패키지 주소 설치 -> 선언

	// 선언 후 go get 설치 예제(엑셀 파일 읽기)
	xfile := "sample.xlsx"

	xlFile, err := xlsx.OpenFile(xfile)

	if err != nil {
		panic("Excel loads error!")
	}

	for _, sheet := range xlFile.Sheets {
		sheet.ForEachRow(func(r *xlsx.Row) error {
			r.ForEachCell(func(c *xlsx.Cell) error {
				text := c.String()
				fmt.Printf("%s\t", text)
				return nil
			})
			return nil
		})

		
	}
}


