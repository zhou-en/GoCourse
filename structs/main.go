package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

func main() {
	en := person{firstName: "En", lastName: "Zhou"}
	fmt.Println(en)

	var alex person
	alex.lastName = "Anderson"
	alex.firstName = "Alex"
	fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@email.com",
			zipCode: 94000,
		},
	}
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson person) print() {
	fmt.Printf("\n%+v", pointerToPerson)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
