package main

import "fmt"

func main() {
	var elements = make([]string, 4)
	elements[0] = "1"
	elements[1] = "2"
	elements[2] = "3"
	elements[3] = "4"
	fmt.Println(elements)

	var newElements = make([]string, 2)
	/*  newElements = append(elements[:0], elements[:2]...)
	fmt.Println(newElements) */

	copy(newElements, elements[2:])
	fmt.Println(newElements)
}
