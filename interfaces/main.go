package main

import "fmt"

//calls all getGreeting() of type string
//calls print greeting because it is of type bot
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//very custom logic for generating an english greeting
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	//very custom logic for generating a spanish greeting
	return "Hola"
}
