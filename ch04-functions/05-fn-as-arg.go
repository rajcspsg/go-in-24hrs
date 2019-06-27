package main

import "fmt"

func anotherFunc(f func() string) string {
	return f()
}

func main() {
	fn := func() string {
		return "func is called"
	}
	fmt.Println(anotherFunc(fn))
}
