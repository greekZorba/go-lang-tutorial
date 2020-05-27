package main

import (
	"fmt"
	"reflect"
)

type car struct {
	name    string "차랑명"
	color   string "색상"
	company string "제조사"
}

func main() {
	// 필드 태그 사용

	tag := reflect.TypeOf(car{})

	for i := 0; i < tag.NumField(); i++ {
		fmt.Println("car : ", tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type)
	}
}
