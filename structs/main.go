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
	// %+v shows the key within the scruct
	// & is an operator. gives access to where alex value is.
	// alexPointer := &alex
	alex.updateName("Alexander")
	alex.print()
}

//* operator followed by memory address, or pointer. pointertoperson is memory address that alex lives at
//* in front of type is different than * in front of pointer.
//*person in a place where type is supposed to be, it is a description of a type. here, looking for a pointer to a person.
//*pointertoperson. operator that takes pointer and assigns to a value. Here, references first name property then updates it to newFirstName

//turn address into value with *address
//turn value into address with &value
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
