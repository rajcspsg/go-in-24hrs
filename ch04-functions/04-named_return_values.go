package main

import "fmt"

func namedreturnValues() (x, y string) {
	x = "raj"
	y = "kumar"
	return
}

func main() {
	fmt.Println(namedreturnValues())
}
