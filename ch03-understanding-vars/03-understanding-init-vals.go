package main

import "fmt"

func main() {
	var i int
	var b bool
	var s string
	var f float32

	fmt.Println("%v %v %v %q", i, f, b, s)

	if s == "" {
		fmt.Println("s is not assigned value yet")
	}

}
