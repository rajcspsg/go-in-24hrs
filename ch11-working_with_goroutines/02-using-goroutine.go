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
	go SlowFunc()
	fmt.Println("I am  shown straight away")
	time.Sleep(time.Second * 2)
}
