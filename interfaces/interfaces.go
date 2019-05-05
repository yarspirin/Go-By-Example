package main

import "fmt"

type Animal interface {
	makeSound()
}

type Cat struct {
}

func (cat Cat) makeSound() {
	fmt.Println("Meow")
}

type Dog struct {
}

func (dog Dog) makeSound() {
	fmt.Println("Woof")
}

func main() {
	dog := Dog{}
	cat := Cat{}

	dog.makeSound()
	cat.makeSound()
}
