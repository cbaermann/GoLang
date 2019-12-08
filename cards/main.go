package main

//:= initializing and assigning values

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")

}

//type conversion example
// greeting := "Hi There!"
// fmt.Println([]byte(greeting))
