package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()} 	//deck in defined as a string slice type in deck.go
	cards = append(cards, "Four of Spades")

	cards.print() 		//print method can be used on cards bcoz the receiver type for
						//print method is deck
}

func newCard() string {
	return "Five of Diamonds"
}
