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
	contact   contactInfo		//embedded struct
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "party",
		contact: contactInfo{			//initializing struct and embedded struct
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}
	fmt.Printf("%+v", jim)
}
