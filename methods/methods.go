package main

import "fmt"

type Item struct {
	price int
	name  string
}

func (item Item) getPrice() int {
	return item.price
}

func (item *Item) changePrice(value int) {
	item.price += value
}

func main() {
	item := Item{2200, "Laptop"}
	fmt.Println(item.getPrice())

	item.changePrice(300)
	fmt.Println(item.getPrice())
}
