package main

//:= initializing and assigning values

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

//type conversion example
// greeting := "Hi There!"
// fmt.Println([]byte(greeting))
