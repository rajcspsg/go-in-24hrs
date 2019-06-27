package main

import "fmt"

func main() {
	s := "Hello World"
	fmt.Println(&s)

	i := 1
	fmt.Println(&i)
	showMemoryAddress(i)
}

func showMemoryAddress(x int) {
	fmt.Println(&x)
	return
}

func showMemoryAddressPointer(x *int) {
	fmt.Println(x)
	return
}
