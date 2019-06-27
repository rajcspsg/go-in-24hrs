package main

import (
	"fmt"
	"time"
)

func pinger(messages chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		chnMsg := <-t.C
		fmt.Println("chnMsg ", chnMsg)
		messages <- "ping"
		chnMsg = <-t.C
		fmt.Println("chnMsg ", chnMsg)
		chnMsg = <-t.C
		fmt.Println("chnMsg ", chnMsg)
	}
}
func main() {
	messages := make(chan string)
	go pinger(messages)

	time.Sleep(5 * time.Second)

	message := <-messages
	fmt.Println(message)

	time.Sleep(5 * time.Second)
}
