package lib

import "fmt"

func init() {
	fmt.Println("lib package Init Start!")
}

func CheckNum(c int32) bool {
	return c > 10
}
