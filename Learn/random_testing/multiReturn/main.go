package main

import "fmt"

func main() {
	cards := newDeck()
	//cards.print()

	deal, remainingcards := deal(cards, 5)
	deal.print()
	fmt.Println("=======")
	remainingcards.print()
}