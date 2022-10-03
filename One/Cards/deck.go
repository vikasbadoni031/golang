//deck.go file
package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {    		// two arguments
	return d[:handSize], d[handSize:]					//returning two objects of type deck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
