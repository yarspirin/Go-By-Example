package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string, 1)

	go func(channel chan string) {
		time.Sleep(time.Second * 4)
		channel <- "Success 1"
	}(chan1)

	select {
	case msg := <-chan1:
		fmt.Println(msg)
	case <-time.After(time.Second * 3):
		fmt.Println("Request timeout")
	}

	chan2 := make(chan string, 1)

	go func(channel chan string) {
		time.Sleep(time.Second * 1)
		channel <- "Success 2"
	}(chan2)

	select {
	case msg := <-chan2:
		fmt.Println(msg)
	case <-time.After(time.Second * 2):
		fmt.Println("Request timeout")
	}
}
