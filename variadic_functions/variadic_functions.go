package main

import "fmt"

func sum(nums ...int) int {
	res := 0

	for _, val := range nums {
		res += val
	}

	return res
}

func main() {
	five := sum(1, 2, 3, 4, 5)
	two := sum(1, 2)

	fmt.Println(five)
	fmt.Println(two)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceSum := sum(slice...)
	fmt.Println(sliceSum)
}
