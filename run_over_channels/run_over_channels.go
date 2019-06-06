package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "Hello, "
	messages <- "world!"

	close(messages)

	for msg := range messages {
		fmt.Print(msg)
	}
	fmt.Println()
}
