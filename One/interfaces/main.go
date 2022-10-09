//file main.go
package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	eb.getGreeting()
	sb.getGreeting()

	printGreeting(eb)
	printGreeting(sb)

}

type bot interface {
	getGreeting() string
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
