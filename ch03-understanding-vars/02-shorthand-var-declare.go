package main

import "fmt"

func main() {
	var s, t string = "foo", "bar"

	var (
		p string = "foo"
		q int    = 2
	)

	fmt.Println(s)
	fmt.Println(t)

	fmt.Println(p)
	fmt.Println(q)
}
