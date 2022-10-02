package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()} //defining a new slice,

	for i, card := range cards { //range is a special keyword that helps loop through the slice.
		fmt.Println(i, card) //we used := bcoz during each iteration the old vars are thrown off.
	}

}

func newCard() string {
	return "Five of Diamonds"
}
