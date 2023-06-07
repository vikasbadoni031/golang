// Testing: calling a func associated with an embedded struct
package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

func (c contactInfo) printContact() {
	fmt.Println("updateEmail called", c.email)
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
	jim.updateName("soemrting")
	jim.print()
	jim.printContact()
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) { //function receving a pointer arg
	(*pointerToPerson).firstName = newFirstName
	(*pointerToPerson).contactInfo.email = "viki.badoni"
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
