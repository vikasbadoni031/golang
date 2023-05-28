package main

import "fmt"

type deck []string


func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuits{
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}


func deal (d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}