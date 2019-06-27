package main

import (
	"fmt"
	"time"
)

func ping1(c chan string) {
	time.Sleep(1 * time.Second)
	c <- "ping from channel1"
}

func ping2(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "ping from channel2"
}
func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go ping1(c1)
	go ping2(c2)

	select {
	case msg1 := <-c1:
		fmt.Println("received", msg1)
	case msg2 := <-c2:
		fmt.Println("received", msg2)
	}

}
