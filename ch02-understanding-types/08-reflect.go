package main

import "fmt"
import "reflect"

func main() {
	var s string = "string"
	var f float32 = 3.2
	var i int32 = 32

	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(i))

}
