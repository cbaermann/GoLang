package main

import "fmt"

func main() {
	card := newCard
	//makes go infer type. only using colon equals syntax when defining a new variable.
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
