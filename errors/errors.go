package main

import "fmt"

type DivisionByZeroError struct {
	a, b int
}

func (error DivisionByZeroError) Error() string {
	return fmt.Sprintf("Can't divide %d by %d", error.a, error.b)
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return -1, DivisionByZeroError{a, b}
	}

	return a / b, nil
}

func main() {
	as := []int{5, 3}
	bs := []int{2, 0}

	for index, a := range as {
		b := bs[index]

		if division, error := divide(a, b); error == nil {
			fmt.Println("Division result", division)
		} else {
			fmt.Println(error)
		}
	}
}
