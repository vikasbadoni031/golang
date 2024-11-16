package main

import (
	"fmt"
	"time"
)

type order struct { //defining a new type
	id        string
	amount    float32
	createdAt time.Time // here we declared the type of  the var "createdAt" to be "Time". and this we did using an external package and the reason for custom type was to later use other funtion provided by the time package on it for eg may be ISO week (https://pkg.go.dev/time#Time.ISOWeek)
	status    string
}

func newOrder(id string, amount float32, status string) *order { //another way to initialize an instance of type "order" //return value is a pointer of type "order"
	myorder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myorder //for any struct created we should return the address as pointer, bcoz else it works on copy and not the oroginal item
}

func main() {
	myorder1 := order{ //created 1st instance of that new type
		id:     "some",
		amount: 23.1,
		status: "Received",
	}
	fmt.Println(myorder1.status) //struct fields can be accessed for get/update the specific element using Dot(.) notation
	fmt.Println(myorder1)
	myorder1.createdAt = time.Now()
	// Important to remember here: we directly called a package and a method inside it using dot notation.
	// as you look at the side doc here. https://pkg.go.dev/time#Time
	// "now" method returns a "Time" type
	// there are also several other methods which has a receiver as type time.
	// which means it can be called on the var we creted above using "createdAt time.Time"

	myorder2 := newOrder("2", 34.2, "Completed") //another way people use to initialise the new instances of a type
	fmt.Println(myorder2)
}
