package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Print(i)

		if i%2 == 0 {
			fmt.Println(" is even")
		} else if i%2 == 1 {
			fmt.Println(" is odd")
		}
	}
}
