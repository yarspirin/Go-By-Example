package main

import "fmt"

func factorize(num int) (int, int) {
	for n1 := 2; n1 <= num; n1++ {
		if num%n1 == 0 {
			return n1, num / n1
		}
	}

	return 1, num
}

func main() {
	num := 77
	p, q := factorize(num)
	fmt.Println(num, "is equal to", p, "x", q)
}
