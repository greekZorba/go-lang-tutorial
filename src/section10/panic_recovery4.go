package main

import (
	"fmt"
	"os"
)

func fileOpen(filename string) {
	// defer 함수(Panic 호출되면 실행)
	defer func() {
		if s := recover(); s != nil {
			fmt.Println("file open error : ", s)
		}
		
	}()

	f, error := os.Open(filename)
	if error != nil {
		panic(error)
	} else {
		fmt.Println("file name : ", f.Name())
	}

	defer f.Close()

}

func main() {
	fileOpen("undefined.txt")
	fmt.Println("End Main !")
}
