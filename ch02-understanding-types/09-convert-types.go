package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var s string = "true"
	//var b bool = true
	b, _ := strconv.ParseBool(s)
	fmt.Println(fmt.Sprintf("str to bool %b", b))

	b = true
	fmt.Println(reflect.TypeOf(b))

	s = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(s))
}
