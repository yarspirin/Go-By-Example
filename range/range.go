package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := []string{"a", "b", "c"}
	m := map[string]int{"RTSI": 150, "S&P 500": 170}
	word := "Hello, World"

	// Iterating through array
	for _, val := range arr {
		fmt.Println(val)
	}

	// Iterating through slice
	for index, val := range slice {
		fmt.Println(index, val)
	}

	// Iterating through map
	for key, value := range m {
		fmt.Println(key, value)
	}

	// Only keys
	for key := range m {
		fmt.Println(key)
	}

	// Through a string
	for char, unicode := range word {
		fmt.Println(char, unicode)
	}
}
