package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// john := person{firstName: "John", lastName: "Walker"}
	// var ted person
	// ted.firstName = "Ted"
	// ted.lastName = "Tomas"

	jim := person{
		firstName: "Jim",
		lastName:  "James",
		contactInfo: contactInfo{
			email:   "jim@jim.com",
			zipCode: 94000,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("Joao")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
