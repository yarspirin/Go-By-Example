package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string, 1)
	chan2 := make(chan string, 2)

	go func(channel chan string) {
		time.Sleep(time.Second * 1)
		channel <- "Hello, "
	}(chan1)

	go func(channel chan string) {
		time.Sleep(time.Second * 2)
		channel <- "World"
	}(chan2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Print(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)
		}
	}
}
