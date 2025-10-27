package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f64 float64
	var str string
	var i32 int32
	var b bool

	var name bool
	fmt.Println(f64, reflect.TypeOf(f64))
	fmt.Println(name, reflect.TypeOf(str))
	fmt.Println(name, reflect.TypeOf(i32))
	fmt.Println(name, reflect.TypeOf(b))

	//위배사항 처음 쓸떄 숫자쓰면안된다
	//대문자 대신 소문자 쓰면 fmt.println 안된다 fmt내에서만 작동
	//단어 바뀔때 대문자로 바꾼다
	// 줄일수 있을때 줄여라

}
