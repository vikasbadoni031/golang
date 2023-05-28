package main

import "fmt"

func main() {

	colors := map[string]string{ //first string for key and second one for values option3 to create map and initialize
		"red":   "#ffakjsbd",
		"green": "#asvbdkjasbd",
		"white": "#kajsbd",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
