package main

import "fmt"

type deck []string //defining a new type

func (d deck) print() { // The receiver here is type deck,
	for i, card := range d { //so this method can be called on any object with type deck
		fmt.Println(i, card) // d is a reference to the cards from main.go
	}
}
