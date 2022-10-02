package main

import "fmt"

type deck []string 

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuits { 				//unusead vars can be replaced by _
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() { 
	for i, card := range d { 
		fmt.Println(i, card) 
	}
}
