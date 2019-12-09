package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Alex", lastName: "Volkanovski"}
	// max := person{firstName: "Max", lastName: "Holloway"}

	fmt.Println(alex)
}
