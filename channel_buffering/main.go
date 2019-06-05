package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	go func() {
		channel <- " world"
	}()

	go func() {
		channel <- "Hello,"
	}()

	message := <-channel
	message += <-channel

	fmt.Println(message)
}
