//main.go file
package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
}
