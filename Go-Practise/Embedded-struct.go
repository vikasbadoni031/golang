package main

import (
	"fmt"
	"time"
)

type customerDetails struct {
	name string
	age  int
}

type order struct {
	id       string
	status   string
	amount   float32
	createAt time.Time
	cdetails customerDetails //we have embedded the customerDetails struct here
}

func main() {
	order1 := order{
		id:     "3",
		status: "Transit",
		amount: 32,
		cdetails: customerDetails{
			name: "Mike",
			age:  25,
		},
	}
	fmt.Println(order1)
	fmt.Println(order1.cdetails.name) //updating the value inside embedded struct
	order1.cdetails.age = 30
	fmt.Println(order1)

}
