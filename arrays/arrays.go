package main

import "fmt"

func main() {
	var a [5]int

	for i := 0; i < 5; i++ {
		a[i] = i + 1
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	fmt.Println(a)

	b := [5]int{5, 6, 6, 1, 2}
	fmt.Println(b)

	var C [5][5]int
	C[0][0] = 1
	for i := 1; i < len(C); i++ {
		C[i][i] = 1
		C[i][0] = 1
		for j := 1; j < len(C[i])-1; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}

	fmt.Println(C)
}
