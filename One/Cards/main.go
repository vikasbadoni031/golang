package main

import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)
}

func newCard() string { //defined the return type of the function
	return "Five of Diamonds"
}
