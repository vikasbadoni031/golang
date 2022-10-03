package main

import "fmt"

type person struct { //defining a struct
	firstName string
	lastName  string
}

func main() {
	//alex := person{"Alex", "Anderson"}
	alex := person{firstName: "Alex", lastName: "Anderson"} // creating an instance of the struct
	fmt.Println(alex)
}
