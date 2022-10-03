//main.go file
package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toStrings())
}
