package main

import (
	"fmt"
)

type contactInfo struct {
	email  string
	mobile int32
}

func (ci contactInfo) contactPrint() {
	fmt.Println("Printing inside contactPrint function")
	fmt.Println(ci.email)
	fmt.Println(ci.mobile)
}

type person struct {
	Name     string
	LastName string
	Age      int32
	contactInfo
}

func main() {
	Sharma := person{
		Name:     "Sharma",
		LastName: "Somoething",
		Age:      30,
		contactInfo: contactInfo{
			email:  "a@vc.com",
			mobile: 989,
		},
	}
	address := &Sharma
	fmt.Printf("%p\n", address)
	fmt.Println(address.email) //you can call the nested structs attribute directly on the address value.
	address.contactPrint()
	// fmt.Println(address.LastName)
}
