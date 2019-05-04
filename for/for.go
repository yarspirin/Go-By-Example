package main

import "fmt"

func main() {
	fmt.Println("for of first type")
	var i = 1
	for i <= 5 {
		fmt.Println(i)
		i += 1
	}
	fmt.Println()

	fmt.Println("for of second(standard) type")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)

	}
	fmt.Println()

	for {
		fmt.Println("Infinite loop")
		fmt.Println()
		break
	}

	for n := 2; n <= 100; n++ {
		var d = 2
		var composite = false

		for d*d <= n {
			if n%d == 0 {
				composite = true
				break
			}

			d += 1
		}

		if !composite {
			fmt.Println(n)
		}
	}
}
