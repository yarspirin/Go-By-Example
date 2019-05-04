package main

import "fmt"

func main() {
	// Empty
	m := make(map[string]int)
	m["Yaroslav"] = 20
	m["Nik"] = 29
	fmt.Println(m)

	// Get
	nik_age := m["Nik"]
	yar_age := m["Yaroslav"]
	fmt.Println(nik_age * yar_age)

	// Len
	fmt.Println(len(m))

	// Delete
	delete(m, "Yaroslav")
	fmt.Println(m)

	// Blanks
	_, was := m["Yaroslav"]
	fmt.Println(was)

	// In-place init
	inplace := map[int]string{4: "foo", 5: "bar"}
	val, has := inplace[6]
	fmt.Println(inplace)
	fmt.Println(val, has)
}
