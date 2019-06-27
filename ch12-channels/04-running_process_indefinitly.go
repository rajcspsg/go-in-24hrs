package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "ping"
		// time.Sleep(5 * time.Second)
		m := <-t.C
		fmt.Println(" pinger time ", m)
	}
}

func main() {
	c := make(chan string)
	go pinger(c)
	for {
		m := <-c
		fmt.Println("main ", m)
	}

}
