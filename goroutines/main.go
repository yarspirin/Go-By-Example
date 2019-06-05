package main

import "fmt"

func f(from string) {
	for i := 0; i < 50; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("directly from main")

	go f("goroutine")
	go f("sug")

	go func(msg string) {
		fmt.Println(msg)
	}("Hello, world!")

	fmt.Println("Done")
}
