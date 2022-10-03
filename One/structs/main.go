package main

import "fmt"

type person struct { //defining a struct
	firstName string
	lastName  string
}

func main() {
	var alex person //when u define a struct bit no values defined
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)				//print fled name and its values
}
