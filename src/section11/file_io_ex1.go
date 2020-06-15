package main

import (
	"fmt"
	"io/ioutil"
	"os"
)
func main() {
	// 파일 읽기, 쓰기 -> ioutil 패키지 사용
	// 더욱 편리하고 직관적으로 파일을 읽고 쓰기 가능
	// 아래 메서드 확인 가능
	// WriteFile(), ReadFile(), ReadAll() 등 사용 가능
	// https://golang.org/pkg/io/ioutil/

	s := "Hello Golang\n File Writes Test!\n"

	// 파일 모드(chmod, chown, chgrp) -> 권한 퍼미션
	// 읽기: 4, 쓰기: 2, 실행: 1
	// 소유자, 그룹, 기타 사용자 순서 (777 : 모든 사용자에게 읽기, 쓰기, 실행 권한부여)
	// https://golang.org/pkg/os/#FileMode
	error := ioutil.WriteFile("text_write1.txt", []byte(s), os.FileMode(0644))
	errorCheck(error)

	data, error := ioutil.ReadFile("sample.txt")
	errorCheck(error)

	fmt.Println("=========================")
	fmt.Println(string(data))
	fmt.Println("=========================")

}

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}
