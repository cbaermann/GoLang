package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Volkanovski",
		contactInfo: contactInfo{
			email:   "alex@tigermuaythai.com",
			zipCode: 83000,
		},
	}
	// %+v shows the key value within the scruct
	alex.updateName("Alexander")
	alex.print()
}
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
