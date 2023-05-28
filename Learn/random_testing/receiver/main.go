package main


func main() {
	cards := deck{"card1", newCard()}
	cards = append(cards, "card3")
	cards.print()
	
}



func newCard() string {
	return "card2"
}