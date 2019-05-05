package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func main() {
	a := 4
	b := 3

	sum := add(a, b)
	mult := multiply(a, b)

	fmt.Println("Sum is ", sum)
	fmt.Println("Multiplication is ", mult)
}
