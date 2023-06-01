package main

import (
	"fmt"
)

func main() {
	authorOne := author{
		name:            "vikas",
		lastName:        "Badoni",
		salary:          30000,
		publish_article: 10,
		total_article:   20,
	}

	scientistOne := scientist{
		name:                "Mohan",
		lastName:            "sharma",
		salary:              40000,
		total_invention:     40,
		published_invention: 25,
	}

	// 1. Directly calling the functions
	// authorOne.PersonDetails()
	// authorOne.PersonProgress()
	// scientistOne.PersonDetails()
	// scientistOne.PersonProgress()

	// 2. Creating var using a common type (nested interface) PersonFinalDetails and calling methods on that.
	/*
		Why interface? as you see any developer interested in the details
		of the person, irrespective of its field (author or scientist)
		doesnt need to go inside the implemnetation details to know what they are doing
		and call a same method PersonDetails/ PersonProgress in both the cases and expect some standard values (eg: string)
	*/

	var f PersonFinalDetails = authorOne
	f.PersonDetails()
	f.PersonProgress()

	var g PersonFinalDetails = scientistOne
	g.PersonDetails()
	g.PersonProgress()

	// This one is not a nested interface.
	// Single level interface used for looping throght structs of various types.

	// anyPerson := []Vikas{authorOne, scientistOne} //single level interface
	// for _, person := range anyPerson {
	// 	person.PersonDetails()
	// 	person.PersonProgress()
	// }

	fmt.Println("another option")
	anyPerson := []PersonFinalDetails{authorOne, scientistOne}
	for _, person := range anyPerson {
		person.PersonDetails()
		person.PersonProgress()
	}

}
