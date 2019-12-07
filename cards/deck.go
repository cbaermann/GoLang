package main

import "fmt"

//create a new type of 'deck'
//which is a slice of strings
type deck []string

//d deck is a reciever. any variable of type deck now gets access to print method.
func (d deck) print() {
	//index, card = current card iterating over, take slice of 'cards' and loop over it
	for i, card := range d {
		fmt.Println(i, card)
	}
}
