package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 1; i <= 100; i++ {
		switch i % 5 {
		case 0:
			fmt.Println("Gives 0")
		case 1:
			fmt.Println("Gives 1")
		case 2:
			fmt.Println("Gives 2")
		case 3:
			fmt.Println("Gives 3")
		case 4:
			fmt.Println("Gives 4")
		default:
			fmt.Println("Unreachable code.")
		}
	}

	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	fmt.Println(time.Now())

	detectType := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("It's a bool")
		case string:
			fmt.Println("It's a string")
		case int:
			fmt.Println("It's an int")
		default:
			fmt.Println("I don't know what it is")
		}
	}

	detectType(4)
	detectType(true)
	detectType("abacada")
	detectType(time.Now())
}
