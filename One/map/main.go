package main

import "fmt"

func main() {
	//var colors map[string]string  //option1 to create map

	colors := make(map[string]string) //option2 to create map

	// colors := map[string]string{ //first string for key and second one for values option3 to create map and initialize
	// 	"red":   "#ffakjsbd",
	// 	"green": "#asvbdkjasbd",
	// }

	colors["red"] = "#jkabsdjk" //adding element to the struct, maps use square brace syntax

	delete(colors, "red")  		//deleting element from a map
	
	fmt.Println(colors)
}
