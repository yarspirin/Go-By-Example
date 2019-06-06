package main

import "fmt"

func main() {
	messages := make(chan string, 1)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("Received msg: ", msg)
	default:
		fmt.Println("No message received.")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("Sent message: ", msg)
	default:
		fmt.Println("No message sent.")
	}

	select {
	case msg := <-messages:
		fmt.Println("Received msg", msg)
	case sig := <-signals:
		fmt.Println("Got signal", sig)
	default:
		fmt.Println("No message received.")
	}
}
