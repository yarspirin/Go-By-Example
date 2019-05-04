package main

import "fmt"

func main() {
	// Creation
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	// Append
	s = append(s, "d")

	// Copy
	var t = make([]string, len(s))
	copy(t, s)

	// Slices
	fmt.Println(s[1:3])

	// 2D

	rows := 5
	columns := 6

	var mat = make([][]int, rows)
	for i := 0; i < rows; i++ {
		mat[i] = make([]int, columns)

		for j := 0; j < columns; j++ {
			mat[i][j] = i*columns + j + 1
		}
	}

	fmt.Println(mat)

	fmt.Println(s)
	fmt.Println(t)
}
