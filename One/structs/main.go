//main.go file
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
	jim := person{
		firstName: "Jim",
		lastName:  "party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	//jimPointer := &jim
	//jimPointer.updateName("Jimmy")
	jim.updateName("jimmy")     // another way to write above two lines, go auto turns the person to "pointer to person"
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {   //function receving a pointer arg
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
