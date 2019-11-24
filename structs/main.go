package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main()  {
	//var alex person
	//
	//alex.firstName = "alex"
	//alex.lastName = "anderson"
	//
	//fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName: "Parson",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}
	//jim.print()

	//jimPointer := &jim  // give memory address of the value
	//jimPointer.updateName("Tom")
	jim.updateName("Tommy")
	jim.print()

}

// function always works on copy of input
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName  // *p get the value from memory address p
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

