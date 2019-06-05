package main

import "fmt"

func doWork(done chan bool) {
	fmt.Println("Counting from 1..100")

	for i := 0; i < 100; i++ {
	}

	fmt.Println("Finished")
	done <- true
}

func main() {
	channel := make(chan bool, 1)
	go doWork(channel)

	<-channel
}
