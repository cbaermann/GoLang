package main

import (
	"fmt"
	"strings"
)

//create a new type of 'deck'
//which is a slice of strings
type deck []string

//slice

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	//use underscore when not using i or j or any var indicator in append/return ect. tells go essentialy that you know you're not using them.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//d deck is a reciever. any variable of type deck now gets access to print method.
func (d deck) print() {
	//index, card = current card iterating over, take slice of 'cards' and loop over it
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
