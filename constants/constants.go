package main

import "fmt"
import "math"

func main() {
	const n = 50000
	const f = 12.3
	const str = "sat"
	const anstr = "3-"

	fmt.Println(n)
	fmt.Println(f)
	fmt.Println(anstr + str)

	var sigmoid = 1 / (1 + math.Exp(-f))
	fmt.Println(sigmoid)
}
