package main

import "fmt"

func trueInc(num *int) int {
	*num++
	return *num
}

func falseInc(num int) int {
	num++
	return num
}

func main() {
	num := 5

	falseInc(num)
	fmt.Println(num) // 5

	trueInc(&num)
	fmt.Println(num) // 6
}
