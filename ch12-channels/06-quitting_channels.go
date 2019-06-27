package main

import (
	"fmt"
	"time"
)

func sender(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "I'm sending message "
		<-t.C
	}
}
func main() {

	messages := make(chan string)
	stop := make(chan bool)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Time's up!")
		stop <- true
	}()
	for {
		select {
		case <-stop:
			return
		case msg := <-messages:
			fmt.Println(msg)
		}
	}
}
