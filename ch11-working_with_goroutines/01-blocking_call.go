package main

import (
	"fmt"
	"time"
)

func SlowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("sleeper() finished")
}
func main() {
	SlowFunc()
	fmt.Println("I am not shown until slowFunc() completes")
}
