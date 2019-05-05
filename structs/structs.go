package main

import "fmt"

type Student struct {
	name    string
	surname string
	age     int
}

func main() {

	yar := Student{"Yar", "Spirin", 20}
	fmt.Println(yar)

	spirin := &yar
	spirin.name = "Nik"
	spirin.age = 29

	fmt.Println(yar)
	fmt.Println(*spirin)

	nik := *spirin
	nik.name = "Nikita"
	yar.name = "Yar"
	fmt.Println(yar)
	fmt.Println(spirin)
	fmt.Println(nik)
}
