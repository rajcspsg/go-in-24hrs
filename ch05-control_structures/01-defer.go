package main

import "fmt"

func main() {
	defer fmt.Println("This will run at the end of the function")
	defer fmt.Println("This will run at the very end of the function")
	defer fmt.Println("This will run at the very very end of the function")
	fmt.Println("Hello World")

}
