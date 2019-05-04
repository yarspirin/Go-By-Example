package main

import "fmt"

func main() {
	var a = "infer a string value"

	b := "shorthand for declaration"
	fmt.Println(a)
	fmt.Println(b)

	var c = a + ", " + b
	fmt.Println(c)

	var d, z = 4, 2
	fmt.Println(d + z)

	var p int
	fmt.Println(p)
}
