package main

import "fmt"

func main() {
	channel := make(chan string)

	go func(msg string) {
		channel <- msg
	}("Hello, world")

	result := <-channel
	fmt.Println(result)
}
