package main

import "fmt"

//:= initializing and assigning values

func main() {
	cards := newDeck()

	fmt.Println(cards.toString())

}

//type conversion example
// greeting := "Hi There!"
// fmt.Println([]byte(greeting))
